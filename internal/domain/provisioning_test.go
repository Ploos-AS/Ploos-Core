package domain

import "testing"

func TestProvisioningRequest(t *testing.T) {
	p := Plan{ID: "p", Product: ProductDomains, ProviderCode: "ploos", Code: "basic"}
	s := Subscription{ID: "s", OrganizationID: "o", PlanID: "p", Status: StatusPending}
	r := ProvisioningRequest{OrganizationID: "o", SubscriptionID: "s", ResourceID: "r", CorrelationID: "c", IdempotencyKey: "k", Product: ProductDomains, Action: ProvisionCreate}
	if err := r.Validate(s, p); err != nil {
		t.Fatal(err)
	}
	for _, action := range []ProvisioningAction{ProvisionSuspend, ProvisionResume, ProvisionDelete, "unknown"} {
		r.Action = action
		if r.Validate(s, p) == nil {
			t.Fatalf("pending subscription accepted %s", action)
		}
	}
	r.Action = ProvisionCreate
	r.OrganizationID = "other"
	if r.Validate(s, p) == nil {
		t.Fatal("cross-tenant request accepted")
	}
	r.OrganizationID = "o"
	r.IdempotencyKey = ""
	if r.Validate(s, p) == nil {
		t.Fatal("missing idempotency key accepted")
	}
	r.IdempotencyKey = "k"
	r.Product = ProductMail
	if r.Validate(s, p) == nil {
		t.Fatal("wrong product accepted")
	}
}
