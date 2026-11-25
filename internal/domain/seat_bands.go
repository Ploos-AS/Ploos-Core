package domain

// SeatBand describes a suggested package size, not a provider entitlement.
type SeatBand string

const (
	SeatSolo     SeatBand = "solo"
	SeatStarter  SeatBand = "starter"
	SeatStandard SeatBand = "standard"
	SeatPlus     SeatBand = "plus"
	SeatPro      SeatBand = "pro"
	SeatBusiness SeatBand = "business"
)

func (b SeatBand) Seats() (int, bool) {
	switch b {
	case SeatSolo:
		return 1, true
	case SeatStarter:
		return 5, true
	case SeatStandard:
		return 10, true
	case SeatPlus:
		return 25, true
	case SeatPro:
		return 50, true
	case SeatBusiness:
		return 100, true
	default:
		return 0, false
	}
}
