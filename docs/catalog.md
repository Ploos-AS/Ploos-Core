# Catalog and subscription policy

The six product codes match the database foundation. Providers and plans must both be explicitly enabled before ordering. External providers remain disabled until commercial and technical readiness is established.

A plan may have a positive seat limit or no seat limit for products that are not seat-based. Suggested organization tiers are 5, 10, 25, 50 and 100; Solo is 1. Provider-specific limits remain possible. No prices or commercial plans are activated by this change.

New subscriptions start pending. Activation requires successful provisioning and applicable billing checks in the future application service. Active subscriptions may be suspended or cancelled; suspended subscriptions may be reactivated or cancelled. Cancelled is terminal. Repeated events must be handled idempotently by the application layer.

The domain functions do not authenticate callers, persist records, enforce quota usage or process payments. Future application services must authorize the organization, load trusted catalog records, and atomically persist subscription and outbox changes.
