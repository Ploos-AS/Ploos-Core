# ADR 0001 — Start as a modular Go monolith

Status: Accepted for the initial foundation (2026-11-24).

## Context

Ploos Core integrates multiple independently operated products, while the business needs low operational complexity and an inexpensive path to automated recurring services.

## Decision

Start with one Go service and one PostgreSQL database, with explicit domain modules and provider adapter interfaces. Use asynchronous provisioning and a transactional outbox when persistence is introduced. Keep product runtimes independent.

Authentication should integrate with an external identity service; Core owns organization memberships and authorization. Billing should integrate with an accounting/payment provider.

## Consequences

Simpler deployment and local development now; modules can be separated later if real scaling needs appear. Product adapters must have contract tests. Core availability must not be required for existing provisioned service traffic.
