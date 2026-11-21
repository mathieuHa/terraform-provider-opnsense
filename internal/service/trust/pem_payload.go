package trust

import (
	"encoding/base64"
	"strings"
)

func pemText(v string) string {
	if v == "" || strings.HasPrefix(strings.TrimSpace(v), "-----BEGIN") {
		return v
	}
	b, err := base64.StdEncoding.DecodeString(strings.TrimSpace(v))
	if err != nil {
		return v
	}
	return string(b)
}
