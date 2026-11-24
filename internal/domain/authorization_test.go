package domain

import "testing"

func TestAuthorize(t *testing.T) {
	for _, role := range []MembershipRole{RoleOwner, RoleAdmin, RoleMember, RoleViewer, "invalid"} {
		m := Membership{OrganizationID: "org-a", UserID: "user-a", Role: role}
		for _, permission := range []Permission{PermissionRead, PermissionManageServices, PermissionManageMembers, PermissionManageOwners, PermissionManageBilling, "invalid"} {
			want := false
			switch permission {
			case PermissionRead:
				want = role != "invalid"
			case PermissionManageServices, PermissionManageMembers:
				want = role == RoleOwner || role == RoleAdmin
			case PermissionManageOwners, PermissionManageBilling:
				want = role == RoleOwner
			}
			if got := Authorize("user-a", "org-a", m, permission); got != want {
				t.Errorf("role %s permission %s: got %v want %v", role, permission, got, want)
			}
			if Authorize("user-a", "org-b", m, permission) || Authorize("user-b", "org-a", m, permission) {
				t.Fatal("cross-tenant or cross-user authorization")
			}
		}
	}
}

func TestMembershipChanges(t *testing.T) {
	admin := Membership{OrganizationID: "org-a", UserID: "user-a", Role: RoleAdmin}
	owner := RoleOwner
	member := RoleMember
	adminRole := RoleAdmin
	if CanChangeMembership("user-a", "org-a", admin, &member, &owner) ||
		CanChangeMembership("user-a", "org-a", admin, &owner, nil) ||
		CanChangeMembership("user-a", "org-a", admin, nil, &adminRole) {
		t.Fatal("admin may not change privileged memberships")
	}
	if !CanChangeMembership("user-a", "org-a", admin, nil, &member) {
		t.Fatal("admin should be able to invite member")
	}
	admin.Role = RoleOwner
	if !CanChangeMembership("user-a", "org-a", admin, &member, &owner) {
		t.Fatal("owner should be able to promote member")
	}
}
