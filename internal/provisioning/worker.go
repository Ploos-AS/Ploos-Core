package provisioning

import (
	"context"
	"errors"
	"time"

	"github.com/Ploos-AS/Ploos-Core/internal/domain"
)

// Job is a leased, persisted outbox request with trusted catalog records.
type Job struct {
	LeaseID      string
	Request      domain.ProvisioningRequest
	Plan         domain.Plan
	Provider     domain.Provider
	Subscription domain.Subscription
}

// Queue must atomically lease jobs and fence completion by lease ID.
// Claim must enforce current desired state, ownership, capacity and retention policy.
type Queue interface {
	Claim(context.Context) (*Job, error)
	Complete(context.Context, string, Result) error
	Retry(context.Context, string, time.Duration) error
}

type Worker struct {
	Queue    Queue
	Registry *Registry
}

// RunOnce processes at most one job. Adapters must be idempotent across retries.
func (w Worker) RunOnce(ctx context.Context) error {
	if w.Queue == nil || w.Registry == nil {
		return errors.New("worker not configured")
	}
	job, err := w.Queue.Claim(ctx)
	if err != nil || job == nil {
		return err
	}
	if job.LeaseID == "" {
		return errors.New("missing lease")
	}
	if err := job.Request.Validate(job.Subscription, job.Plan); err != nil {
		return w.Queue.Retry(ctx, job.LeaseID, time.Minute)
	}
	if job.Request.Action == domain.ProvisionCreate && !job.Plan.CanOrder(job.Provider) {
		return w.Queue.Retry(ctx, job.LeaseID, time.Minute)
	}
	adapter, err := w.Registry.Resolve(job.Plan, job.Provider)
	if err != nil {
		return w.Queue.Retry(ctx, job.LeaseID, time.Minute)
	}
	result, err := adapter.Execute(ctx, job.Request)
	if err != nil {
		return w.Queue.Retry(ctx, job.LeaseID, time.Minute)
	}
	return w.Queue.Complete(ctx, job.LeaseID, result)
}
