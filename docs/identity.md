# Identity boundary

Core identifies users by verified issuer and subject, never email or client-provided internal IDs. The middleware requires a verifier and resolver and fails closed when absent. The storage lookup does not automatically create accounts. Organization authorization remains separate.

This is the authentication boundary only, not an OIDC implementation. A concrete access-token verifier, issuer and audience configuration, JWKS rotation, resolver wiring, onboarding and live provider tests are required before mounting customer endpoints. Browser login and session management are separate. Authelia is a candidate, not a mandatory provider.
