# PostgreSQL migration runner

Use PostgreSQL 16 and the psql client. Run with a dedicated migration role and a configured libpq service (PGSERVICE), not the runtime application's credentials.

    psql -X -v ON_ERROR_STOP=1 -f migrations/apply.psql

The runner serializes migration sessions with an advisory lock and records the foundation version atomically with its schema changes. Repeated runs skip applied migrations. Existing manually initialized databases require an explicit audited baseline; do not automatically mark them migrated.

This is an operator-side migration runner, not a Go database connection pool. Runtime credentials, readiness checks and database-backed repositories are follow-up work. Do not run the legacy down migration against production data.
