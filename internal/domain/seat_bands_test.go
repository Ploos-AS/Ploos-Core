package domain

import "testing"

func TestSeatBands(t *testing.T) {
	for band, want := range map[SeatBand]int{
		SeatSolo: 1, SeatStarter: 5, SeatStandard: 10,
		SeatPlus: 25, SeatPro: 50, SeatBusiness: 100,
	} {
		got, ok := band.Seats()
		if !ok || got != want {
			t.Fatalf("%s: got %d, want %d", band, got, want)
		}
	}
	for _, band := range []SeatBand{"", "enterprise", "unknown"} {
		if _, ok := band.Seats(); ok {
			t.Fatalf("unexpected fixed band %q", band)
		}
	}
}
