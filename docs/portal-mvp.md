# Portal API pilot

The Core executable can explicitly enable authenticated read-only portal routes using PLOOS_CORE_PORTAL_ENABLED=true, a configured PostgreSQL database, and the following environment variables:

- PLOOS_CORE_AUTH_ENDPOINT
- PLOOS_CORE_AUTH_ISSUER
- PLOOS_CORE_AUTH_AUDIENCE
- PLOOS_CORE_AUTH_CLIENT_ID
- PLOOS_CORE_AUTH_SECRET

Keep secrets in a protected service environment. Default is disabled. Verify provider access-token introspection compatibility before enabling even a pilot. No browser frontend, login flow, organization discovery, checkout or write operations are included. Existing users and memberships must be provisioned through trusted administration. The API is not production-ready.
