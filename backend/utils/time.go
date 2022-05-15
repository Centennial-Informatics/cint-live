package utils

import (
	"time"
)

func IsBefore(t time.Time) bool {
	return time.Until(t) >= 0
}

func IsAfter(t time.Time) bool {
	return time.Since(t) >= 0
}

/* Use case: do something 15mins after contest ends. */
func XAfter(x time.Duration, t time.Time) time.Time {
	return t.Add(x)
}

func XBefore(x time.Duration, t time.Time) time.Time {
	return t.Add(-x)
}
