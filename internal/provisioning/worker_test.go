package provisioning

import (
	"context"
	"testing"
	"time"

	"github.com/Ploos-AS/Ploos-Core/internal/domain"
)

type queueStub struct {
	job       *Job
	completed bool
	retried   bool
}

func (q *queueStub) Claim(context.Context) (*Job, error) { return q.job, nil }
func (q *queueStub) Complete(context.Context, string, Result) error {
	q.completed = true
	return nil
}
func (q *queueStub) Retry(context.Context, string, time.Duration) error {
	q.retried = true
	return nil
}

func TestWorker(t *testing.T) {
	registry := NewRegistry()
	if err := registry.Register(Route{Product: domain.ProductCloud, Provider: "ploos"}, stubAdapter{}); err != nil {
		t.Fatal(err)
	}
	for _, enabled := range []bool{true, false} {
		q := &queueStub{job: &Job{
			LeaseID:      "lease",
			Request:      domain.ProvisioningRequest{OrganizationID: "o", SubscriptionID: "s", ResourceID: "r", CorrelationID: "c", IdempotencyKey: "k", Product: domain.ProductCloud, Action: domain.ProvisionCreate},
			Plan:         domain.Plan{ID: "p", Product: domain.ProductCloud, ProviderCode: "ploos", Code: "starter", Enabled: enabled},
			Provider:     domain.Provider{Code: "ploos", Enabled: true},
			Subscription: domain.Subscription{ID: "s", OrganizationID: "o", PlanID: "p", Status: domain.StatusPending},
		}}
		if err := (Worker{Queue: q, Registry: registry}).RunOnce(context.Background()); err != nil {
			t.Fatal(err)
		}
		if q.completed != enabled || q.retried == enabled {
			t.Fatal("unexpected job outcome")
		}
	}
}
