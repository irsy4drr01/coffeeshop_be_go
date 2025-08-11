package utils

import (
	"time"
)

// Hanya format tanggal, tanpa embel-embel "since"
func FormatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("02 January 2006")
}
