# Organization storage

Organization creation and initial ownership are atomic. Membership mutations load the actor's current role from storage and lock the organization row before reading memberships. All membership writers must use this same locking protocol; direct SQL writes can bypass the invariant. The final owner cannot be removed or demoted.

Actor and initial owner IDs must be resolved from verified identities by a future application service. No public HTTP endpoints are introduced. User onboarding, audit events, invitations, concurrent integration tests and runtime role grants remain follow-up work. PostgreSQL CI runs the integration tests using its disposable database.
