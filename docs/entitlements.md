# Entitlements and quotas

Capacity is granted per organization, subscription and quota key. Examples: seats, storage_bytes, mailboxes, websites and domains. Missing grants deny new allocation; zero grants no capacity. A five-seat package grants five seats, not five organization administrators. Organization memberships and product accounts are separate concepts.

Only active subscriptions can allocate new capacity. This policy does not delete existing resources on suspension or downgrade. Usage must come from trusted product records, never from client input. Concurrent allocation requires atomic reservation and reconciliation; this PR implements only the pure domain check. Provider-specific quota units and shared mailbox rules must be defined by adapters.

Suggested seat tiers remain 1, 5, 10, 25, 50 and 100, without hardcoding them as the only allowed limits. No commercial packages are activated.
