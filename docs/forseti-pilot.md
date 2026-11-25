# Forseti pilot preparation

After merging, manually run Publish pilot image on main. Record the published digest and use ghcr.io/ploos-as/ploos-core@sha256:<digest> in the Compose environment. Do not deploy from a mutable tag.

Use the existing Forseti directory standard and a dedicated Compose project. Inspect existing listening ports before selecting a loopback port. Use an approved PostgreSQL 16 image digest and protected environment/password files. Start only db first, apply migrations from the same Core revision, create a dedicated login role ploos_core_runtime with a strong password using a protected operator session, then apply deploy/pilot-runtime-grants.sql. Never put passwords in Git or shell history.

The initial runtime role is read-only and must be a new role without inherited privileges. Configure the Core DSN using that role, database host db, and keep both portal and organization creation disabled. Start Core, check /healthz and /readyz locally, inspect logs, and verify that PostgreSQL has no published port. Test backup and restore before adding any customer data. Do not expose Caddy or DNS routes yet.

This is preparation only. The operator must verify the host, available port, registry access, database permissions and volume placement. No production deployment is authorized by this document.
