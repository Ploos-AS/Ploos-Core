package domain

import "testing"

func TestOrganizationValidation(t *testing.T) {
	tests := []struct {
		name string
		org  Organization
		ok   bool
	}{
		{"association", Organization{ID: "org-1", DisplayName: "Example association", Kind: KindAssociation}, true},
		{"business", Organization{ID: "org-2", DisplayName: "Example company", Kind: KindBusiness}, true},
		{"missing id", Organization{DisplayName: "Example", Kind: KindBusiness}, false},
		{"missing name", Organization{ID: "org-3", Kind: KindBusiness}, false},
		{"unknown kind", Organization{ID: "org-4", DisplayName: "Example", Kind: "unknown"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.org.Validate()
			if (err == nil) != tt.ok {
				t.Fatalf("Validate() error=%v, want valid=%v", err, tt.ok)
			}
		})
	}
}

func TestMembershipValidation(t *testing.T) {
	valid := Membership{OrganizationID: "org-1", UserID: "user-1", Role: RoleAdmin}
	if err := valid.Validate(); err != nil {
		t.Fatal(err)
	}
	valid.Role = "superuser"
	if err := valid.Validate(); err == nil {
		t.Fatal("expected error for unknown role")
	}
}
