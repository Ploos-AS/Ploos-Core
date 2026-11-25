# Identity and portal next milestone

Proposed hostnames: auth.ploos.no for identity, account.ploos.no for the customer portal, api.ploos.no for Core. Pilot hostnames and public exposure require separate approval.

## Identity provider evaluation

Evaluate Authelia, Authentik, Keycloak and Zitadel against OIDC authorization code with PKCE, MFA/passkeys, account recovery, customer self-registration, access-token audience and introspection compatibility, revocation, operational cost, licensing and backup/restore. Verify actual token claims and reject ID tokens at the Core API. No provider is selected yet.

Core owns organization membership and product entitlements. The provider owns authentication. Do not infer membership from email or identity provider groups.

## Portal MVP

Use a separate Ploos-Portal repository. Start with a mock-data frontend, then connect to the authenticated Core API after provider verification. Required views: organization selector, organization list, product catalog, subscription overview and account settings. Prefer Norwegian and English from the outset. Never embed provider secrets or expose access tokens in logs. Evaluate a server-side session/BFF with secure cookies and CSRF protection before live integration.

Ordering, payments, provisioning and organization creation remain disabled until their authorization and operational prerequisites are satisfied. Keep the existing Forseti Core pilot private.

## Acceptance gates

Verify identity provider interoperability, explicit enrollment, tenant isolation, logout/session expiry and recovery in a disposable environment before connecting the portal to the pilot. Complete backup automation and private monitoring before external customer access.
