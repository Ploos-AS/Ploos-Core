BEGIN;
DELETE FROM core_schema_migrations WHERE version = 3;
DROP TABLE IF EXISTS domains_organization_links;
DELETE FROM core_schema_migrations WHERE version = 2;
DROP TABLE audit_events, provisioning_outbox, entitlements, subscriptions, plans, providers, products, memberships, users, organizations;
DELETE FROM core_schema_migrations WHERE version = 1;
COMMIT;
