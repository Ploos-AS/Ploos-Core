# Pilot database backup and restore

Run the reviewed deploy/backup-pilot.sh as root on Forseti. It creates a protected PostgreSQL custom-format dump, checks its archive listing and atomically renames the completed file. No automatic deletion or scheduling is configured.

A listing check is not a restore test. To test recovery, use a separate disposable PostgreSQL 16 instance with no published ports and a fresh database. Create the required runtime role separately, then restore with pg_restore --exit-on-error --no-owner --no-acl into the fresh database. Verify schema versions 1 through 4 and reapply the reviewed pilot-runtime-grants.sql. Never restore over the live database for a test.

The dump contains database data and schema, not cluster roles, runtime secrets, Compose configuration or image digests. Protect and back up these separately. Copy encrypted backups off-host and verify recovery before enabling customer access. Retention, encryption keys, recovery objectives and automation remain operator decisions.
