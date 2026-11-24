# Provider adapters

Route by product and provider, using trusted subscription and catalog records. The registry is populated at startup; unknown routes fail closed.

Planned routes include Domains/Ploos, DNS/Ploos, Hosting/Ploos, Mail/Ploos, Cloud/Ploos, Cloud/Infomaniak, Cloud/Proton and Forenio/Ploos. No production adapter is registered by this change.

Provider and plan enablement must be checked before new allocations. Routing itself deliberately permits disabled providers so existing subscriptions can still be reconciled or safely decommissioned. A worker must enforce action-specific readiness, authorization, ownership, retention, idempotency and desired-state checks before execution. Adapter credentials must remain server-side.

Nextcloud region selection belongs to trusted plan configuration; Norwegian and European hosting can use separate plans without separate customer identities.

This PR adds an interface and registry with unit tests only. No network calls, live provisioning, retries or external credentials are introduced.
