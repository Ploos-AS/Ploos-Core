# Organization creation pilot

POST /api/v1/organizations accepts display_name and kind (business, association, individual). The server generates a UUID and derives the initial owner from verified identity. This creates a Core tenant, not a legally registered entity. Creation is separately gated by PLOOS_CORE_ORGANIZATION_CREATE_ENABLED=true. Keep disabled outside controlled pilots. Rate limiting, idempotency, verified legal authority and integration tests remain required before public signup.
