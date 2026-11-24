package domain

import (
	"math"
	"testing"
)

func TestAllocation(t *testing.T) {
	e := Entitlement{OrganizationID: "a", SubscriptionID: "s", Key: "seats", Limit: 5}
	s := Subscription{ID: "s", OrganizationID: "a", Status: StatusActive}
	for _, tc := range []struct {
		org, key        string
		used, requested int64
		ok              bool
	}{
		{"a", "seats", 4, 1, true},
		{"a", "seats", 5, 1, false},
		{"a", "seats", 0, 5, true},
		{"a", "seats", 0, 6, false},
		{"b", "seats", 0, 1, false},
		{"a", "storage", 0, 1, false},
		{"a", "seats", -1, 1, false},
		{"a", "seats", 0, 0, false},
		{"a", "seats", 1, math.MaxInt64, false},
	} {
		if got := e.CheckAllocation(s, tc.org, tc.key, tc.used, tc.requested) == nil; got != tc.ok {
			t.Errorf("allocation %+v: got %v", tc, got)
		}
	}
	s.Status = StatusSuspended
	if e.CheckAllocation(s, "a", "seats", 0, 1) == nil {
		t.Fatal("suspended subscription accepted")
	}
	s.Status = StatusActive
	s.ID = "other"
	if e.CheckAllocation(s, "a", "seats", 0, 1) == nil {
		t.Fatal("wrong subscription accepted")
	}
}
