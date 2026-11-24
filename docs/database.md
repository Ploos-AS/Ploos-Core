# Database foundation

Target PostgreSQL 16+. Apply to a fresh development database using:
```sh
psql "$DATABASE_URL" -v ON_ERROR_STOP=1 -f migrations/0001_foundation.up.sql
```

The down migration destroys all foundation data; use only in disposable environments. Production migration tracking, backup and recovery must be implemented before deployment.

External identities are keyed by issuer and subject, not email. Composite foreign keys prevent entitlements and outbox events from referencing another organization's subscription. This does not replace request authorization or database access policies; those are follow-up work.

All providers and plans default to disabled. No reseller agreement or commercial offer is implied. Suggested seat bands are 1, 5, 10, 25, 50 and 100; plans may differ by provider and product. Seat limits are nullable for non-seat products such as domains. Entitlements store actual subscription quotas separately from catalog defaults.

Subscription and outbox writes must share a transaction. Worker leases, retries and delivery implementation are deferred. Audit events intentionally omit arbitrary payloads and secrets. No customer-facing API is enabled by this migration.
