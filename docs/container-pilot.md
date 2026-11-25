# Container pilot

Build locally with docker build -t ploos-core:pilot . . The runtime is non-root and has a liveness health check. Readiness is separately exposed at /readyz and requires schema version 4.

The Compose template requires explicit images, an unused loopback port, protected runtime environment file and database password file. PostgreSQL is not published to the host. Set PLOOS_CORE_ADDR=0.0.0.0:8080 in the runtime environment if overriding the image default. Keep portal and organization creation disabled initially.

Before starting Core, initialize PostgreSQL, create a separate least-privilege runtime role, apply migrations/apply.psql using a privileged operator session, and grant only necessary runtime permissions. Use the runtime role in PLOOS_CORE_DATABASE_URL, with db as the host. Do not use the PostgreSQL superuser in Core. No automatic migration is performed.

For Forseti use /etc/ploos-apps/ploos-core/compose.yaml and /srv/ploos-apps/ploos-core/{data,releases}. The template uses a named database volume; backup and volume placement must be explicitly reviewed. Choose and verify the loopback port before installation. Do not add Caddy or public DNS until readiness and access controls are verified.

Ploos Deploy remains a separate execution boundary; this template is not a validated Deploy manifest. Base image tags are not digest-pinned yet; pin approved images before production. No live deployment or GHCR publication is performed by this change.
