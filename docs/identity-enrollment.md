# Central identity enrollment

Core identifies accounts by exact verified issuer and subject, not email. One internal user can hold memberships in multiple organizations. Enrollment is an explicit trusted operator operation, separate from authentication and membership grants.

EnrollIdentity is an internal storage primitive, not an HTTP endpoint or CLI. Its caller must independently verify the identity and approved issuer. Repeated enrollment returns the existing internal identifier. Different issuers remain separate even when subjects match. No organization membership is granted.

The read-only pilot runtime role cannot enroll users. Keep portal disabled until a real identity provider is tested for access-token introspection, audience, issuer, expiry and revocation. Authelia remains a candidate, not a verified integration. Operator authorization, audit, account recovery, invitation flow and provider selection remain pending. Never link accounts automatically by email.
