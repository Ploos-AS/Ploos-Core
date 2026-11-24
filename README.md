# Ploos Core

Ploos Core is the shared, open-source control plane for Ploos AS products.

It connects **Ploos Domains, DNS, Hosting, Mail, Cloud** and **Forenio** using common organization identities, memberships, product catalogs, subscriptions and provisioning requests. Product applications remain independently deployable and operational.

## Principles

- **Multi-tenant from day one:** organizations own subscriptions and resources; people may belong to several organizations.
- **One identity, multiple organizations:** never equate a login with a tenant.
- **Provider neutral:** Ploos-operated infrastructure and reseller integrations use the same adapter boundary.
- **Least privilege:** scope each action to an organization, actor and explicit permission; never trust an organization ID supplied by a client alone.
- **Safe provisioning:** requests use stable idempotency keys and asynchronous delivery; workers must tolerate retries and out-of-order events.
- **Data portability:** provide export and deletion processes with documented retention rules.
- **Product isolation:** customer websites, DNS, mail and cloud services must keep running if Core is unavailable.

## Scope

Core owns the **organization registry, membership/roles, product catalog, entitlements, subscription state, provisioning orchestration and audit trail**. It does not operate DNS nameservers, web servers, mail servers, Nextcloud or Forenio itself.

Identity authentication is an external integration, with Core responsible for identity-to-organization authorization. Billing and payments should be integrated behind an adapter, not implemented as a payment processor.

## Initial package layout

- `internal/domain/`: shared organization and entitlement types and validation.
- `docs/architecture.md`: system boundaries, flows, and rollout plan.
- `docs/decisions/0001-modular-monolith.md`: initial architecture decision.

## Development

Go 1.24+ is the initial implementation language. Run `go test ./...` locally.

This is a foundation, **not** a production-ready customer portal, payment service or identity provider.
