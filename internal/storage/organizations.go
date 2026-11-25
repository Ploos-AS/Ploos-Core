package storage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Ploos-AS/Ploos-Core/internal/domain"
)

// CreateOrganization atomically creates a tenant and its initial owner.
// ownerID must be resolved from a verified identity by the caller.
func CreateOrganization(ctx context.Context, db *sql.DB, org domain.Organization, ownerID string) error {
	if err := org.Validate(); err != nil {
		return err
	}
	m := domain.Membership{OrganizationID: org.ID, UserID: ownerID, Role: domain.RoleOwner}
	if err := m.Validate(); err != nil {
		return err
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if _, err := tx.ExecContext(ctx, "INSERT INTO organizations(id, display_name, kind) VALUES ($1,$2,$3)", org.ID, org.DisplayName, org.Kind); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "INSERT INTO memberships(organization_id,user_id,role) VALUES ($1,$2,$3)", org.ID, ownerID, domain.RoleOwner); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "INSERT INTO audit_events(organization_id,actor_user_id,action,resource_id) VALUES ($1,$2,$3,$4)", org.ID, ownerID, "organization.created", org.ID); err != nil {
		return err
	}
	return tx.Commit()
}

// ChangeMembership serializes membership changes per tenant and preserves
// at least one owner. actorID must come from a verified identity.
func ChangeMembership(ctx context.Context, db *sql.DB, orgID, actorID, targetID string, proposed *domain.MembershipRole) error {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var locked string
	if err := tx.QueryRowContext(ctx, "SELECT id FROM organizations WHERE id=$1 FOR UPDATE", orgID).Scan(&locked); err != nil {
		return err
	}
	actor := domain.Membership{OrganizationID: orgID, UserID: actorID}
	if err := tx.QueryRowContext(ctx, "SELECT role FROM memberships WHERE organization_id=$1 AND user_id=$2", orgID, actorID).Scan(&actor.Role); err != nil {
		return err
	}
	var role domain.MembershipRole
	var current *domain.MembershipRole
	err = tx.QueryRowContext(ctx, "SELECT role FROM memberships WHERE organization_id=$1 AND user_id=$2", orgID, targetID).Scan(&role)
	if err == nil {
		current = &role
	} else if !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if !domain.CanChangeMembership(actorID, orgID, actor, current, proposed) {
		return errors.New("membership change denied")
	}
	if current != nil && *current == domain.RoleOwner && (proposed == nil || *proposed != domain.RoleOwner) {
		var count int
		if err := tx.QueryRowContext(ctx, "SELECT count(*) FROM memberships WHERE organization_id=$1 AND role='owner'", orgID).Scan(&count); err != nil {
			return err
		}
		if count <= 1 {
			return errors.New("last owner cannot be removed")
		}
	}
	if proposed == nil {
		_, err = tx.ExecContext(ctx, "DELETE FROM memberships WHERE organization_id=$1 AND user_id=$2", orgID, targetID)
	} else {
		_, err = tx.ExecContext(ctx, "INSERT INTO memberships(organization_id,user_id,role) VALUES ($1,$2,$3) ON CONFLICT (organization_id,user_id) DO UPDATE SET role=EXCLUDED.role", orgID, targetID, *proposed)
	}
	if err != nil {
		return err
	}
	return tx.Commit()
}
