// Package gigasecond provides methods to calculate when a person has
// lived for one gigasecond (10^9 seconds).
package gigasecond

import (
	"time"
)

const testVersion = 4

// AddGigasecond returns the time when a person has lived for one gigasecond.
func AddGigasecond(b time.Time) time.Time {
	return b.Add(1E18)
}
