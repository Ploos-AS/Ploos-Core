BEGIN;
CREATE TABLE domains_organization_links (
 organization_id uuid PRIMARY KEY REFERENCES organizations(id),
 domains_organization_id uuid NOT NULL UNIQUE,
 created_at timestamptz NOT NULL DEFAULT now()
);
INSERT INTO core_schema_migrations(version) VALUES (2);
COMMIT;
