package domain

import (
	"errors"
	"strings"
)

type ProductCode string

const (
	ProductDomains ProductCode = "domains"
	ProductDNS     ProductCode = "dns"
	ProductHosting ProductCode = "hosting"
	ProductMail    ProductCode = "mail"
	ProductCloud   ProductCode = "cloud"
	ProductForenio ProductCode = "forenio"
)

func (p ProductCode) Valid() bool {
	switch p {
	case ProductDomains, ProductDNS, ProductHosting, ProductMail, ProductCloud, ProductForenio:
		return true
	default:
		return false
	}
}

type Provider struct {
	Code    string
	Enabled bool
}

type Plan struct {
	ID           string
	Product      ProductCode
	ProviderCode string
	Code         string
	SeatLimit    *int
	Enabled      bool
}

func (p Plan) Validate() error {
	if strings.TrimSpace(p.ID) == "" || strings.TrimSpace(p.Code) == "" || strings.TrimSpace(p.ProviderCode) == "" || !p.Product.Valid() {
		return errors.New("invalid plan identity")
	}
	if p.SeatLimit != nil && *p.SeatLimit <= 0 {
		return errors.New("seat limit must be positive")
	}
	return nil
}

// CanOrder requires both the plan and its matching provider to be enabled.
func (p Plan) CanOrder(provider Provider) bool {
	return p.Validate() == nil && p.Enabled && provider.Enabled && p.ProviderCode == provider.Code
}

type SubscriptionStatus string

const (
	StatusPending   SubscriptionStatus = "pending"
	StatusActive    SubscriptionStatus = "active"
	StatusSuspended SubscriptionStatus = "suspended"
	StatusCancelled SubscriptionStatus = "cancelled"
)

type Subscription struct {
	ID             string
	OrganizationID string
	PlanID         string
	Status         SubscriptionStatus
}

func NewSubscription(id, organizationID string, plan Plan, provider Provider) (Subscription, error) {
	if strings.TrimSpace(id) == "" || strings.TrimSpace(organizationID) == "" || !plan.CanOrder(provider) {
		return Subscription{}, errors.New("subscription cannot be ordered")
	}
	return Subscription{ID: id, OrganizationID: organizationID, PlanID: plan.ID, Status: StatusPending}, nil
}

// CanTransition defines lifecycle policy, not authorization or billing proof.
func (s SubscriptionStatus) CanTransition(next SubscriptionStatus) bool {
	switch s {
	case StatusPending:
		return next == StatusActive || next == StatusCancelled
	case StatusActive:
		return next == StatusSuspended || next == StatusCancelled
	case StatusSuspended:
		return next == StatusActive || next == StatusCancelled
	default:
		return false
	}
}
