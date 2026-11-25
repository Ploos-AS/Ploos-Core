package cloud

import "testing"

func TestProvider(t *testing.T) {
	for track, want := range map[Track]string{Office: "infomaniak", Privacy: "proton", Drive: "ploos"} {
		got, ok := track.Provider()
		if !ok || got != want {
			t.Fatalf("unexpected provider for %s", track)
		}
	}
	if _, ok := Track("unknown").Provider(); ok {
		t.Fatal("unknown track accepted")
	}
}
