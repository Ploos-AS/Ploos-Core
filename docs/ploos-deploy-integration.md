# Ploos Deploy integration boundary

The Forseti pilot is an existing stateful Compose project named ploos-core, with configuration at /etc/ploos-apps/ploos-core/compose.yaml and protected environment files in the same directory. Core binds to loopback port 18090. PostgreSQL uses the existing project named volume core-db.

Ploos Deploy currently documents an experimental guarded disposable Balder deployment. Do not use its deploy or update commands against Core until explicit support for adopting existing stateful projects is implemented and tested.

## Required adoption contract

- Inspect existing project, service names, image digests, volume identity and configuration without printing secrets.
- Adopt without recreating the database or replacing its volume.
- Require an explicit approved Core digest and a compatible schema version.
- Take and verify a backup before any schema-changing operation.
- Keep migrations as a separate operator-approved operation; application rollback does not imply schema rollback.
- Update only the core service, preserving database and secrets.
- Verify both /healthz and /readyz before recording success.
- Serialize updates and retain previous image digest and deployment metadata.
- Fail closed on unexpected configuration, missing backup or incompatible schema.

Test adoption and failed updates on a disposable stateful project first. No live Core update, Deploy manifest or automatic migration is authorized by this document.
