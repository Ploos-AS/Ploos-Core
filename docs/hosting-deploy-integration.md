# Hosting and Deploy integration boundary

Core owns customer organizations, subscriptions, authorization and provisioning intent. Hosting defines independently purchasable products. Deploy owns local deployment execution and manifests.

The upstream Hosting README describes a planning repository. The Deploy README describes an experimental CLI with guarded test deployment and no public API. Core therefore does not execute SSH, shell commands, Docker operations or Deploy manifests.

For future web hosting, ProjectName derives a stable opaque Compose-compatible project identifier from trusted organization and resource IDs. This is a proposed naming convention, not an existing Deploy contract. Existing workloads must not be renamed or adopted automatically. A persistent resource-to-project mapping, server assignment, ownership checks, approved manifest templates and a narrowly privileged local executor are required before activation.

Mail, SMTP, DNS, storage and workspace remain independent services; a web deployment must not automatically create or cancel them. No live adapter is registered and no production deployment is attempted.
