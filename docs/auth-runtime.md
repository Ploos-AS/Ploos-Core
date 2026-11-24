# Authentication runtime boundary

The configuration builder validates a complete trusted introspection configuration. DatabaseResolver links verified identities to existing Core users. Neither component is exposed through public routes yet.

Before enabling authentication in the executable, verify the provider's access-token-only introspection behavior and required claims, configure secrets through a protected secret source, and run live provider tests. No automatic user creation or email-based account linking is permitted. Authelia compatibility remains unverified.
