package auth

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Introspector delegates access-token validation to a trusted OAuth server.
type Introspector struct {
	endpoint, issuer, audience, clientID, secret string
	client                                     *http.Client
}

func NewIntrospector(endpoint, issuer, audience, clientID, secret string) (*Introspector, error) {
	for _, raw := range []string{endpoint, issuer} {
		u, err := url.Parse(raw)
		if err != nil || u.Scheme != "https" || u.Host == "" || u.User != nil || u.RawQuery != "" || u.Fragment != "" {
			return nil, errors.New("invalid authentication URL")
		}
	}
	if audience == "" || clientID == "" || secret == "" {
		return nil, errors.New("missing authentication configuration")
	}
	return &Introspector{endpoint: endpoint, issuer: issuer, audience: audience, clientID: clientID, secret: secret, client: &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(*http.Request, []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}}, nil
}

func (v *Introspector) Verify(ctx context.Context, token string) (Identity, error) {
	denied := errors.New("access token rejected")
	if strings.TrimSpace(token) == "" {
		return Identity{}, denied
	}
	form := url.Values{"token": {token}, "token_type_hint": {"access_token"}}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, v.endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return Identity{}, denied
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(url.QueryEscape(v.clientID), url.QueryEscape(v.secret))
	resp, err := v.client.Do(req)
	if err != nil {
		return Identity{}, denied
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Identity{}, denied
	}
	const limit = 65536
	body, err := io.ReadAll(io.LimitReader(resp.Body, limit+1))
	if err != nil || len(body) > limit {
		return Identity{}, denied
	}
	var claims struct {
		Active    bool            `json:"active"`
		Issuer    string          `json:"iss"`
		Subject   string          `json:"sub"`
		Audience  json.RawMessage `json:"aud"`
		TokenType string          `json:"token_type"`
		Expires   int64           `json:"exp"`
		NotBefore int64           `json:"nbf"`
	}
	now := time.Now().Unix()
	if json.Unmarshal(body, &claims) != nil || !claims.Active || claims.Issuer != v.issuer || strings.TrimSpace(claims.Subject) == "" || !strings.EqualFold(claims.TokenType, "Bearer") || claims.Expires <= now || claims.NotBefore > now {
		return Identity{}, denied
	}
	var audiences []string
	if json.Unmarshal(claims.Audience, &audiences) != nil {
		var single string
		if json.Unmarshal(claims.Audience, &single) != nil {
			return Identity{}, denied
		}
		audiences = []string{single}
	}
	for _, audience := range audiences {
		if audience == v.audience {
			return Identity{Issuer: claims.Issuer, Subject: claims.Subject}, nil
		}
	}
	return Identity{}, denied
}

var _ Verifier = (*Introspector)(nil)
