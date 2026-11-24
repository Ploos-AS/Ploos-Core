# Organization authorization

Core uses an authenticated external issuer/subject identity mapped to an internal user ID. Each request must resolve the requested organization and load the actor's membership from trusted storage. The domain policy is deny-by-default.

| Permission | Owner | Admin | Member | Viewer |
|---|---|---|---|---|
| Read organization metadata | yes | yes | yes | yes |
| Manage services | yes | yes | no | no |
| Manage ordinary memberships | yes | yes | no | no |
| Change owner/admin memberships | yes | no | no | no |
| Manage billing | yes | no | no | no |

Read permission is not access to private mail or cloud files. Product-level permissions remain independent.

Membership writes must validate the target membership's organization, serialize changes by locking the organization row, and prevent removal or demotion of the final owner in the same transaction. Organization creation and initial owner creation must be atomic. The current PR implements policy functions and tests only, not database-backed enforcement, authentication, or HTTP endpoints.
