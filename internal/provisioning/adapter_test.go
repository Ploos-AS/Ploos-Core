package provisioning

import (
	"context"
	"testing"

	"github.com/Ploos-AS/Ploos-Core/internal/domain"
)

type stubAdapter struct{}

func (stubAdapter) Execute(context.Context, domain.ProvisioningRequest) (Result, error) {
	return Result{ExternalResourceID: "test-resource"}, nil
}

func TestRegistry(t *testing.T) {
	r := NewRegistry()
	route := Route{Product: domain.ProductCloud, Provider: "ploos"}
	if err := r.Register(route, stubAdapter{}); err != nil {
		t.Fatal(err)
	}
	if r.Register(route, stubAdapter{}) == nil {
		t.Fatal("duplicate route accepted")
	}
	p := domain.Plan{ID: "p", Product: domain.ProductCloud, ProviderCode: "ploos", Code: "starter"}
	if _, err := r.Resolve(p, domain.Provider{Code: "ploos"}); err != nil {
		t.Fatal(err)
	}
	if _, err := r.Resolve(p, domain.Provider{Code: "proton"}); err == nil {
		t.Fatal("mismatched provider accepted")
	}
	p.Product = domain.ProductMail
	if _, err := r.Resolve(p, domain.Provider{Code: "ploos"}); err == nil {
		t.Fatal("unregistered product accepted")
	}
}
