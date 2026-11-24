# Runtime database readiness

Set PLOOS_CORE_DATABASE_URL using a secret environment file or secret manager. Use a dedicated runtime role with SELECT access to core_schema_migrations; do not use migration credentials. Require verified TLS for remote databases (sslmode=verify-full). Without a DSN readiness returns 503. With a DSN readiness checks schema version 1 with a two-second deadline. Liveness is independent of the database.

The pool defaults to ten open and two idle connections. No customer API is enabled yet. Dependency checksums were sourced from postgres_exporter and must pass Go module verification in CI.
