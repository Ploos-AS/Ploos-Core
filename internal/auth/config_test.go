package auth

import "testing"

func TestConfigBuild(t *testing.T) {
	c := Config{Endpoint: "https://issuer.example/introspect", Issuer: "https://issuer.example", Audience: "core", ClientID: "core", Secret: "secret"}
	if _, err := c.Build(); err != nil {
		t.Fatal(err)
	}
	c.Secret = ""
	if _, err := c.Build(); err == nil {
		t.Fatal("missing secret accepted")
	}
	c.Secret = "secret"
	c.Endpoint = "http://issuer.example/introspect"
	if _, err := c.Build(); err == nil {
		t.Fatal("insecure endpoint accepted")
	}
}
