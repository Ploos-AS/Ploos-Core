# Provisioning contract

Core issues asynchronous requests scoped to an organization, subscription, resource and product. Each request carries a correlation ID and stable idempotency key.

Create is allowed for pending or active subscriptions; suspend for suspended subscriptions; resume for active subscriptions; delete for cancelled subscriptions. These are necessary conditions, not sufficient authorization. Deletion additionally requires explicit retention-policy and authorization checks. Cancellation must not automatically delete customer data.

The application must authenticate the actor, authorize the operation, verify resource ownership, check provider readiness for new allocations, reserve capacity atomically and write the outbox event in the same database transaction. Repeated keys must return the original outcome or reject mismatched payloads. Workers must reconcile stale events against desired state and support retries without duplicate resources.

This change adds the domain contract only. HTTP endpoints, database-backed orchestration, worker leases, adapter implementations and destructive-action approval are not implemented.
