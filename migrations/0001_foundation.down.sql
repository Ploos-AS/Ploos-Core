BEGIN;
DROP TABLE audit_events, provisioning_outbox, entitlements, subscriptions, plans, providers, products, memberships, users, organizations;
DELETE FROM core_schema_migrations WHERE version = 1;
COMMIT;
