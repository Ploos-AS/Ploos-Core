package domain

import (
	"errors"
	"strings"
)

// Organization is a customer tenancy boundary, not an authenticated user.
type Organization struct {
	ID          string
	DisplayName string
	Kind        OrganizationKind
}

type OrganizationKind string

const (
	KindBusiness    OrganizationKind = "business"
	KindAssociation OrganizationKind = "association"
	KindIndividual  OrganizationKind = "individual"
)

func (o Organization) Validate() error {
	if strings.TrimSpace(o.ID) == "" {
		return errors.New("organization id is required")
	}
	if strings.TrimSpace(o.DisplayName) == "" {
		return errors.New("organization display name is required")
	}
	switch o.Kind {
	case KindBusiness, KindAssociation, KindIndividual:
		return nil
	default:
		return errors.New("unsupported organization kind")
	}
}

// Membership separates a global user identity from access to a tenant.
type Membership struct {
	OrganizationID string
	UserID         string
	Role           MembershipRole
}

type MembershipRole string

const (
	RoleOwner  MembershipRole = "owner"
	RoleAdmin  MembershipRole = "admin"
	RoleMember MembershipRole = "member"
	RoleViewer MembershipRole = "viewer"
)

func (m Membership) Validate() error {
	if strings.TrimSpace(m.OrganizationID) == "" || strings.TrimSpace(m.UserID) == "" {
		return errors.New("membership requires organization and user")
	}
	switch m.Role {
	case RoleOwner, RoleAdmin, RoleMember, RoleViewer:
		return nil
	default:
		return errors.New("unsupported membership role")
	}
}
