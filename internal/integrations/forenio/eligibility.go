package forenio

import "github.com/Ploos-AS/Ploos-Core/internal/domain"

// CanManageServices checks eligibility for optional association services.
// The free formation wizard must never call this gate.
func CanManageServices(org domain.Organization, actorID string, membership domain.Membership) bool {
	return org.Validate() == nil && org.Kind == domain.KindAssociation &&
		domain.Authorize(actorID, org.ID, membership, domain.PermissionManageServices)
}
