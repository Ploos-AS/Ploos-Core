package domain

import "testing"

func TestPlanOrdering(t *testing.T) {
	seats := 5
	p := Plan{ID: "p", Product: ProductCloud, ProviderCode: "ploos", Code: "starter", SeatLimit: &seats, Enabled: true}
	provider := Provider{Code: "ploos", Enabled: true}
	s, err := NewSubscription("s", "org", p, provider)
	if err != nil || s.Status != StatusPending || s.OrganizationID != "org" || s.PlanID != "p" {
		t.Fatalf("unexpected subscription: %+v, %v", s, err)
	}
	for _, other := range []Provider{{Code: "ploos"}, {Code: "proton", Enabled: true}} {
		if p.CanOrder(other) {
			t.Fatal("disabled or mismatched provider accepted")
		}
	}
	p.Enabled = false
	if p.CanOrder(provider) {
		t.Fatal("disabled plan accepted")
	}
	p.Enabled = true
	seats = 0
	if p.CanOrder(provider) {
		t.Fatal("invalid seat limit accepted")
	}
	p.SeatLimit = nil
	p.Product = ProductDomains
	if !p.CanOrder(provider) {
		t.Fatal("non-seat plan rejected")
	}
	p.Product = "unknown"
	if p.CanOrder(provider) {
		t.Fatal("unknown product accepted")
	}
}

func TestSubscriptionTransitions(t *testing.T) {
	statuses := []SubscriptionStatus{StatusPending, StatusActive, StatusSuspended, StatusCancelled, "unknown"}
	allowed := map[SubscriptionStatus][]SubscriptionStatus{
		StatusPending:   {StatusActive, StatusCancelled},
		StatusActive:    {StatusSuspended, StatusCancelled},
		StatusSuspended: {StatusActive, StatusCancelled},
	}
	for _, from := range statuses {
		for _, to := range statuses {
			want := false
			for _, candidate := range allowed[from] {
				want = want || candidate == to
			}
			if from.CanTransition(to) != want {
				t.Errorf("unexpected transition %s -> %s", from, to)
			}
		}
	}
}
