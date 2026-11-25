BEGIN;
ALTER TABLE provisioning_outbox
 ADD COLUMN lease_id uuid UNIQUE,
 ADD COLUMN lease_until timestamptz,
 ADD COLUMN dead_at timestamptz,
 ADD CONSTRAINT outbox_lease_pair CHECK ((lease_id IS NULL) = (lease_until IS NULL));
INSERT INTO core_schema_migrations(version) VALUES (3);
COMMIT;
