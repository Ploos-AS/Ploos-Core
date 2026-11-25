\set ON_ERROR_STOP on
-- Run as database owner after migrations, in the dedicated ploos_core database.
-- Create the LOGIN role ploos_core_runtime separately with a protected password.
BEGIN;
REVOKE CREATE ON SCHEMA public FROM PUBLIC;
GRANT USAGE ON SCHEMA public TO ploos_core_runtime;
GRANT SELECT ON core_schema_migrations, users, organizations, memberships, products, providers, plans, subscriptions TO ploos_core_runtime;
-- Initial pilot is read-only. Organization creation stays disabled.
COMMIT;
