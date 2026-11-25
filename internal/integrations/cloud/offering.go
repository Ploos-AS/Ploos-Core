package cloud

// Track identifies the planned Cloud service family, not an active subscription.
type Track string

const (
	Office  Track = "office"
	Privacy Track = "privacy"
	Drive   Track = "drive"
)

// Provider returns the expected Core catalog provider for a known track.
func (t Track) Provider() (string, bool) {
	switch t {
	case Office:
		return "infomaniak", true
	case Privacy:
		return "proton", true
	case Drive:
		return "ploos", true
	default:
		return "", false
	}
}
