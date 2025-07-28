package overlap

import (
	"time"
)

type DateRange struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// Overlaps returns true if two date ranges overlap.
func Overlaps(a, b DateRange) bool {
	return !a.End.Before(b.Start) && !b.End.Before(a.Start)
}
