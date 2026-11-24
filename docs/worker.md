# Provisioning worker contract

RunOnce claims at most one leased outbox job, validates its subscription and catalog scope, resolves the adapter, and records completion or schedules a retry. New allocations require enabled plans and providers. Adapters must be idempotent.

The queue interface is not a PostgreSQL implementation. Atomic leases, fencing, bounded retries, dead-letter handling, resource persistence, desired-state reconciliation and retention approval must be implemented before running a worker. The current retry delay is a provisional one minute. No live adapter or worker process is enabled.
