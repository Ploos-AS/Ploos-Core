package domains

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

// Client reads portfolios using a delegated user token, not a global admin token.
type Client struct {
	base string
	http *http.Client
}

func NewClient(base string) (*Client, error) {
	u, err := url.Parse(base)
	if err != nil || u.Scheme != "https" || u.Host == "" || u.User != nil || u.RawQuery != "" || u.Fragment != "" {
		return nil, errors.New("invalid HTTPS Domains base URL")
	}
	return &Client{base: strings.TrimRight(base, "/"), http: &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(*http.Request, []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}}, nil
}

// Portfolio returns opaque domain records until the upstream field contract
// is versioned. organizationID must come from a trusted organization mapping.
func (c *Client) Portfolio(ctx context.Context, organizationID, token string) ([]json.RawMessage, error) {
	if strings.TrimSpace(organizationID) == "" || strings.ContainsAny(organizationID, "/\\") || organizationID == "." || organizationID == ".." || strings.TrimSpace(token) == "" {
		return nil, errors.New("missing or invalid portfolio credentials")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.base+"/api/v1/organizations/"+url.PathEscape(organizationID)+"/portfolio", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, errors.New("Domains request failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Domains portfolio unavailable")
	}
	const maxBody = 2 << 20
	body, err := io.ReadAll(io.LimitReader(resp.Body, maxBody+1))
	if err != nil || len(body) > maxBody {
		return nil, errors.New("invalid Domains response size")
	}
	var result struct {
		Domains []json.RawMessage `json:"domains"`
	}
	if json.Unmarshal(body, &result) != nil || result.Domains == nil {
		return nil, errors.New("invalid Domains portfolio response")
	}
	return result.Domains, nil
}
