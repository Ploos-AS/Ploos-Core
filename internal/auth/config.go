package auth

import "errors"

// Config is supplied by trusted startup configuration, never request data.
type Config struct {
	Endpoint string
	Issuer   string
	Audience string
	ClientID string
	Secret   string
}

func (c Config) Build() (Verifier, error) {
	if c.Endpoint == "" || c.Issuer == "" || c.Audience == "" || c.ClientID == "" || c.Secret == "" {
		return nil, errors.New("incomplete authentication configuration")
	}
	return NewIntrospector(c.Endpoint, c.Issuer, c.Audience, c.ClientID, c.Secret)
}
