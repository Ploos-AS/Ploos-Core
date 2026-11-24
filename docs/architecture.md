# Ploos Core: architecture and integration boundaries

## System map

```text
Ploos Portal / Forenio / approved API clients
                    |
              Ploos Core API
                    |
 Organizations -- Memberships -- Catalog -- Subscriptions
                    |                    |
               Entitlements          Provisioning queue
                                         |
                   Provider adapters / product APIs
                     Domains | DNS | Hosting | Mail | Cloud
                                         |
                                   Ploos Deploy
```

**Forenio remains a distinct product and brand.** A Forenio association may be linked to one Core organization, with explicit permissions and consent. Neither service should bypass the other's authorization checks.

## Domain boundaries

| Capability | Core responsibility | Other system responsibility |
| --- | --- | --- |
| Identity | Map external identity subject to organization memberships and roles | Authenticate users, MFA, sessions |
| Organizations | Stable tenant identifier, metadata and membership | Customer-specific business workflows |
| Domains | Order / entitlement / provisioning orchestration | Registrar integrations and domain lifecycle |
| DNS | Subscription and permissions | Authoritative zones, nameservers, DNSSEC |
| Hosting | Subscription and provisioning request | Websites, databases, backups, deployment |
| Mail | Seat quotas and provision requests | Mailboxes, routing, spam, transport, mail data |
| Cloud | Seat/storage entitlement and provider selection | Nextcloud or approved reseller integrations |
| Forenio | Linked customer organization and purchased product state | Association founding and governance |
| Billing | Subscription state and invoice references | Accounting, taxes, payment collection |

## Critical rules

1. A user can be a member of multiple organizations. Switching organizations must be explicit and checked on every API request.
2. Entitlements are scoped to organization + product + subscription. Seat plans begin with the suggested Solo/1, Starter/5, Standard/10, Plus/25, Pro/50, Business/100 bands; supplier-specific terms may differ.
3. Domain and DNS administration must work independently, including externally registered domains.
4. A product may use an independent control plane. Core is the order and access orchestrator, not the runtime of all products.
5. Provider adapters must not assume a reseller agreement exists (particularly Infomaniak and Proton). Unsupported offers cannot be activated.
6. Provisioning events must include organization ID, resource ID, correlation ID and idempotency key. Use transactional outbox on database-backed implementation.
7. Organization administrators must not automatically gain access to private mailbox contents or private cloud files.
8. Customer-facing service availability cannot depend on synchronous calls to Core.
9. Audit security-sensitive state changes without recording passwords, tokens or mailbox contents.

## Initial implementation milestones

1. Typed domain model, validation and CI.
2. PostgreSQL schema and migrations for organizations, users, memberships, products, subscriptions, entitlements and outbox.
3. Authorization tests for tenant isolation and membership role changes.
4. Versioned HTTP API and OpenAPI contract with explicit error responses.
5. Product adapter contract, mock adapter, idempotent job processor and retry handling.
6. One end-to-end sandbox integration with an existing Ploos product.
7. Billing connector and reconciliation; then portal workflows.

Do not expose a production checkout until provisioning, reconciliation, access control, backup and operational recovery are tested.
