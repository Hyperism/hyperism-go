package utix

import (
	"strings"
)

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
