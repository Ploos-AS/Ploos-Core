BEGIN;
CREATE TABLE organizations (
 id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
 display_name text NOT NULL CHECK (btrim(display_name) <> ''),
 kind text NOT NULL CHECK (kind IN ('business','association','individual')),
 created_at timestamptz NOT NULL DEFAULT now()
);
CREATE TABLE users (
 id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
 issuer text NOT NULL CHECK (btrim(issuer) <> ''),
 subject text NOT NULL CHECK (btrim(subject) <> ''),
 created_at timestamptz NOT NULL DEFAULT now(),
 UNIQUE (issuer, subject)
);
CREATE TABLE memberships (
 organization_id uuid NOT NULL REFERENCES organizations(id),
 user_id uuid NOT NULL REFERENCES users(id),
 role text NOT NULL CHECK (role IN ('owner','admin','member','viewer')),
 PRIMARY KEY (organization_id,user_id)
);
CREATE INDEX memberships_user_idx ON memberships(user_id);
CREATE TABLE products (
 code text PRIMARY KEY,
 display_name text NOT NULL
);
INSERT INTO products VALUES
 ('domains','Ploos Domains'),('dns','Ploos DNS'),('hosting','Ploos Hosting'),
 ('mail','Ploos Mail'),('cloud','Ploos Cloud'),('forenio','Forenio');
CREATE TABLE providers (
 code text PRIMARY KEY,
 display_name text NOT NULL,
 enabled boolean NOT NULL DEFAULT false
);
INSERT INTO providers(code,display_name) VALUES
 ('ploos','Ploos'),('infomaniak','Infomaniak'),('proton','Proton');
CREATE TABLE plans (
 id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
 product_code text NOT NULL REFERENCES products(code),
 provider_code text NOT NULL REFERENCES providers(code),
 code text NOT NULL,
 seat_limit integer CHECK (seat_limit > 0),
 enabled boolean NOT NULL DEFAULT false,
 UNIQUE(product_code,provider_code,code)
);
CREATE TABLE subscriptions (
 id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
 organization_id uuid NOT NULL REFERENCES organizations(id),
 plan_id uuid NOT NULL REFERENCES plans(id),
 status text NOT NULL DEFAULT 'pending' CHECK (status IN ('pending','active','suspended','cancelled')),
 created_at timestamptz NOT NULL DEFAULT now(),
 UNIQUE(organization_id,id)
);
CREATE INDEX subscriptions_org_idx ON subscriptions(organization_id);
CREATE TABLE entitlements (
 organization_id uuid NOT NULL,
 subscription_id uuid NOT NULL,
 quota_key text NOT NULL CHECK (btrim(quota_key) <> ''),
 quota_value bigint NOT NULL CHECK (quota_value >= 0),
 PRIMARY KEY (organization_id,subscription_id,quota_key),
 FOREIGN KEY (organization_id,subscription_id) REFERENCES subscriptions(organization_id,id)
);
CREATE TABLE provisioning_outbox (
 id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
 organization_id uuid NOT NULL,
 subscription_id uuid NOT NULL,
 resource_id text NOT NULL CHECK (btrim(resource_id) <> ''),
 correlation_id uuid NOT NULL,
 idempotency_key text NOT NULL CHECK (btrim(idempotency_key) <> ''),
 event_type text NOT NULL,
 payload jsonb NOT NULL DEFAULT '{}'::jsonb CHECK (jsonb_typeof(payload) = 'object'),
 attempts integer NOT NULL DEFAULT 0 CHECK (attempts >= 0),
 available_at timestamptz NOT NULL DEFAULT now(),
 delivered_at timestamptz,
 created_at timestamptz NOT NULL DEFAULT now(),
 UNIQUE(organization_id,idempotency_key),
 FOREIGN KEY (organization_id,subscription_id) REFERENCES subscriptions(organization_id,id)
);
CREATE INDEX outbox_pending_idx ON provisioning_outbox(available_at,id) WHERE delivered_at IS NULL;
CREATE TABLE audit_events (
 id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
 organization_id uuid NOT NULL REFERENCES organizations(id),
 actor_user_id uuid REFERENCES users(id),
 action text NOT NULL,
 resource_id text NOT NULL,
 created_at timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX audit_org_idx ON audit_events(organization_id,created_at);
CREATE TABLE IF NOT EXISTS core_schema_migrations (
 version integer PRIMARY KEY,
 applied_at timestamptz NOT NULL DEFAULT now()
);
INSERT INTO core_schema_migrations(version) VALUES (1);
COMMIT;
