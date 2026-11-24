package domain

import "errors"

// Entitlement is a subscription-scoped capacity grant.
type Entitlement struct {
	OrganizationID string
	SubscriptionID string
	Key            string
	Limit          int64
}

// CheckAllocation checks a proposed increase against trusted usage.
// Callers must serialize usage checks and reservations in storage.
func (e Entitlement) CheckAllocation(s Subscription, organizationID, key string, used, requested int64) error {
	if organizationID == "" || s.ID == "" || key == "" ||
		e.OrganizationID != organizationID || s.OrganizationID != organizationID ||
		e.SubscriptionID != s.ID || e.Key != key {
		return errors.New("entitlement scope mismatch")
	}
	if s.Status != StatusActive {
		return errors.New("subscription is not active")
	}
	if e.Limit < 0 || used < 0 || requested <= 0 {
		return errors.New("invalid capacity")
	}
	if used > e.Limit || requested > e.Limit-used {
		return errors.New("quota exceeded")
	}
	return nil
}
