# Outbox completion

Schema 4 stores the adapter's external resource identifier on the delivered event. Completion is a single fenced SQL update; expired or superseded tokens cannot complete an event.

This is an event-level result, not a canonical resource inventory. Adapters must return a nonempty identifier even for deletion (the identifier of the deleted resource). A dedicated inventory, provider uniqueness rules, lease renewal, authorization checks and database integration tests remain necessary. No worker is activated.
