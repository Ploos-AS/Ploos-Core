# Ploos Domains integration

The upstream contract is documented in Ploos-AS/Ploos-Domains, docs/DOMAIN_PORTFOLIO_API.md. The current endpoint is GET /api/v1/organizations/{organizationID}/portfolio and returns {"domains":[...]}. It requires configured OIDC verification, trusted identity resolution and an authorized portfolio reader.

Core's initial client is read-only, HTTPS-only, bounded by timeout and response size, and does not follow redirects. It accepts a delegated user token valid for Domains and an upstream organization ID obtained from a trusted mapping. Core and Domains IDs must not be assumed identical. No tokens are persisted by this client.

An empty result may mean no domains or no upstream membership; never interpret it as proof of authorization. Core must independently enforce tenant access before invoking the client. OIDC audience/token exchange, organization linking, pagination, field-level contract and end-to-end tests remain open. The upstream endpoint is not production-ready according to its documentation.

This is not a provisioning adapter: no documented upstream write contract was established for this integration. No registration, DNS mutation, billing or production connection is enabled.
