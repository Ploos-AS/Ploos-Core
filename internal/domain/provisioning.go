package domain

import (
	"errors"
	"strings"
)

type ProvisioningAction string

const (
	ProvisionCreate  ProvisioningAction = "create"
	ProvisionSuspend ProvisioningAction = "suspend"
	ProvisionResume  ProvisioningAction = "resume"
	ProvisionDelete  ProvisioningAction = "delete"
)

type ProvisioningRequest struct {
	OrganizationID string
	SubscriptionID string
	ResourceID     string
	CorrelationID  string
	IdempotencyKey string
	Product        ProductCode
	Action         ProvisioningAction
}

// Validate checks request scope and lifecycle. The application must load
// subscription and plan from trusted storage and authorize the actor.
func (r ProvisioningRequest) Validate(s Subscription, p Plan) error {
	for _, value := range []string{r.OrganizationID, r.SubscriptionID, r.ResourceID, r.CorrelationID, r.IdempotencyKey} {
		if strings.TrimSpace(value) == "" {
			return errors.New("missing provisioning identifier")
		}
	}
	if r.OrganizationID != s.OrganizationID || r.SubscriptionID != s.ID || s.PlanID != p.ID || p.Validate() != nil || r.Product != p.Product {
		return errors.New("provisioning scope mismatch")
	}
	switch r.Action {
	case ProvisionCreate:
		if s.Status == StatusPending || s.Status == StatusActive {
			return nil
		}
	case ProvisionSuspend:
		if s.Status == StatusSuspended {
			return nil
		}
	case ProvisionResume:
		if s.Status == StatusActive {
			return nil
		}
	case ProvisionDelete:
		if s.Status == StatusCancelled {
			return nil
		}
	}
	return errors.New("invalid provisioning action or subscription status")
}
