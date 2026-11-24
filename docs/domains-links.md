# Domains organization links

Migration 2 introduces a one-to-one mapping between Core and Domains organization IDs. The read lookup requires current Core membership. Domains must independently authorize the delegated user token.

Mappings are provisioned only by a trusted operator after verifying ownership on both systems. No public mapping mutation endpoint is provided. Automated linking, audit, revocation coordination and live integration tests remain pending. One-to-one mapping is an initial policy, not a claim that all future organizational structures must be identical.
