# Durable outbox leases

Schema 3 adds lease tokens, expiry and dead-letter timestamps. Claim is atomic, increments attempts, and reserves an event for two minutes. Retry is fenced by the current unexpired token. Exhausted expired events can be marked dead by the reaper.

Defaults are ten attempts and one minute retry delay. The legacy transaction-only helpers must not be mixed with this lease protocol. No worker is activated. Completion with atomic external resource persistence, lease renewal, authorization and desired-state checks, concurrency tests and operator reconciliation are still required. External effects must be idempotent; leases alone do not guarantee exactly-once execution.
