# Runtime database readiness

The storage package provides a database/sql schema check for migration version 1. The HTTP handler accepts a readiness callback with a two-second context deadline. Without a callback readiness remains unavailable.

The PostgreSQL driver, connection pool configuration and main executable wiring are intentionally deferred until dependency checksums can be generated and verified. No database credentials are required or accepted by the executable yet. Runtime integration tests against PostgreSQL remain necessary.
