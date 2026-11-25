# Ploos Core roadmap

Status baseline: after PR 27. Implemented means code exists, not production acceptance.

## Current baseline

- Organization and membership models, PostgreSQL schema 4, identity mapping and role checks.
- Opt-in authenticated catalog and subscription read API.
- Outbox lease, retry, dead-letter and completion primitives; no runnable production queue.
- Domains organization mapping and read client; Cloud, Hosting/Deploy and Forenio integration boundaries, not live provisioning.

## Pilot delivery sequence

1. Organization discovery: authenticated membership-scoped list.
2. Organization creation: verified existing identity, atomic owner assignment, audit and validation.
3. Catalog: disabled-by-default proposed seat bands Solo 1, Starter 5, Standard 10, Plus 25, Pro 50, Business 100; supplier terms override. No prices or resale promises.
4. Security and integration tests: real PostgreSQL tenant isolation, concurrent membership changes, migration upgrades, outbox fencing and expiry, provider access-token verification, malformed requests and HTTP contracts.
5. Container packaging: non-root runtime, pinned dependencies, secrets outside images, separate migration/runtime database roles, health checks.
6. Isolated Forseti pilot: explicit approval, backups and recovery test, no live customer data or external provisioning.
7. Portal frontend: organization selection, catalog and subscriptions; verified login and session design.

## Before commercial launch

- Complete resource inventory, atomic capacity reservations, queue authorization, lease renewal, reconciliation and idempotent sandbox adapter.
- Add ordering, billing adapter, tax/invoice handling, cancellation and retention policies, audit and export.
- Confirm supplier agreements, hosting locations, service levels, DNS/mail operational readiness and unit economics.
- Ensure Core outages cannot interrupt existing customer services.

## Acceptance gates

Do not enable production checkout or live provisioning until security, recovery, billing and reconciliation are tested. Do not enable provider-backed products without contractual and technical readiness. Keep Forenio's free formation flow independent and accountless.

This roadmap is a planning baseline, not a completed code audit. PR numbering is intentionally not a release commitment.
