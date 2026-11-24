package provisioning

import (
	"context"
	"errors"

	"github.com/Ploos-AS/Ploos-Core/internal/domain"
)

// Adapter executes a previously authorized and persisted provisioning request.
// Implementations must support idempotent retries and verify resource ownership.
type Adapter interface {
	Execute(context.Context, domain.ProvisioningRequest) (Result, error)
}

type Result struct {
	ExternalResourceID string
}

type Route struct {
	Product  domain.ProductCode
	Provider string
}

// Registry is configured at startup and must not be mutated after publication.
type Registry struct {
	adapters map[Route]Adapter
}

func NewRegistry() *Registry {
	return &Registry{adapters: make(map[Route]Adapter)}
}

func (r *Registry) Register(route Route, adapter Adapter) error {
	if !route.Product.Valid() || route.Provider == "" || adapter == nil {
		return errors.New("invalid adapter registration")
	}
	if _, exists := r.adapters[route]; exists {
		return errors.New("adapter already registered")
	}
	r.adapters[route] = adapter
	return nil
}

// Resolve uses trusted catalog data, never a provider selected by the client.
func (r *Registry) Resolve(plan domain.Plan, provider domain.Provider) (Adapter, error) {
	if plan.Validate() != nil || plan.ProviderCode != provider.Code {
		return nil, errors.New("invalid provider route")
	}
	adapter, exists := r.adapters[Route{Product: plan.Product, Provider: provider.Code}]
	if !exists {
		return nil, errors.New("adapter not configured")
	}
	return adapter, nil
}
