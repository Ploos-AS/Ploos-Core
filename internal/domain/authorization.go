package domain

// Permission describes an organization-scoped operation.
type Permission string

const (
	PermissionRead           Permission = "read"
	PermissionManageServices Permission = "manage_services"
	PermissionManageMembers  Permission = "manage_members"
	PermissionManageOwners   Permission = "manage_owners"
	PermissionManageBilling  Permission = "manage_billing"
)

// Authorize requires a membership loaded from trusted storage, never from
// client-supplied role claims. Unknown roles and permissions fail closed.
func Authorize(actorID, organizationID string, membership Membership, permission Permission) bool {
	if membership.Validate() != nil || membership.UserID != actorID || membership.OrganizationID != organizationID {
		return false
	}
	switch permission {
	case PermissionRead:
		return true
	case PermissionManageServices, PermissionManageMembers:
		return membership.Role == RoleOwner || membership.Role == RoleAdmin
	case PermissionManageOwners, PermissionManageBilling:
		return membership.Role == RoleOwner
	default:
		return false
	}
}

// CanChangeMembership checks both the current and proposed roles.
// A nil current or proposed role represents creation or removal respectively.
// The caller must enforce the last-owner invariant in a database transaction.
func CanChangeMembership(actorID, organizationID string, actor Membership, current, proposed *MembershipRole) bool {
	if !Authorize(actorID, organizationID, actor, PermissionManageMembers) {
		return false
	}
	if current == nil && proposed == nil {
		return false
	}
	for _, role := range []*MembershipRole{current, proposed} {
		if role == nil {
			continue
		}
		switch *role {
		case RoleOwner, RoleAdmin:
			if !Authorize(actorID, organizationID, actor, PermissionManageOwners) {
				return false
			}
		case RoleMember, RoleViewer:
		default:
			return false
		}
	}
	return true
}
