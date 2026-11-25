package hosting

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

// ProjectName derives an opaque, stable deployment name from trusted Core IDs.
// It is not an authorization check or a Deploy manifest.
func ProjectName(organizationID, resourceID string) (string, error) {
	if strings.TrimSpace(organizationID) == "" || strings.TrimSpace(resourceID) == "" {
		return "", errors.New("missing deployment scope")
	}
	h := sha256.New()
	for _, id := range []string{organizationID, resourceID} {
		if strings.ContainsRune(id, 0) {
			return "", errors.New("invalid deployment scope")
		}
		h.Write([]byte(id))
		h.Write([]byte{0})
	}
	return "ploos-web-" + hex.EncodeToString(h.Sum(nil)[:16]), nil
}
