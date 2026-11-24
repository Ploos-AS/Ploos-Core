# Access token verification

Core API authentication uses a dedicated OAuth 2.0 introspection adapter (RFC 7662), not an OIDC ID-token verifier. The Domains ID-token verifier is intended for a different token class. The configured HTTPS authorization server must validate access tokens and return active, iss, sub, aud, token_type and exp. Core additionally checks issuer, audience, expiry and optional nbf. The endpoint must reject ID tokens and refresh tokens; token_type_hint alone does not enforce that property.

This strict response profile may not be supported by every provider. Authelia compatibility has not been established. No discovery, JWT/JWKS verification, runtime wiring or browser login is enabled. Provider compatibility and negative tests must precede deployment. Secrets are not logged and redirects are disabled.
