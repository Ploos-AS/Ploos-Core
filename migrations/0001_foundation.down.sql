BEGIN;
DROP TABLE audit_events, provisioning_outbox, entitlements, subscriptions, plans, providers, products, memberships, users, organizations;
COMMIT;
