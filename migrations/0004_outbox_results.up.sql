BEGIN;
ALTER TABLE provisioning_outbox ADD COLUMN external_resource_id text;
INSERT INTO core_schema_migrations(version) VALUES (4);
COMMIT;
