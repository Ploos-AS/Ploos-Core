# PostgreSQL outbox locking

LockNextOutbox selects a due event using FOR UPDATE SKIP LOCKED in a caller-owned transaction. No available event returns sql.ErrNoRows. DeferLockedOutbox increments attempts and moves availability forward one minute.

These are transaction primitives, not a complete Queue implementation. A row lock is not a durable lease. Do not execute network calls inside the transaction or treat an unlocked event as owned. A subsequent migration must add durable lease tokens and expiration, fenced completion, retry limits and dead-letter state. No worker is enabled by this change.

Concurrency and database integration tests remain required before production use.
