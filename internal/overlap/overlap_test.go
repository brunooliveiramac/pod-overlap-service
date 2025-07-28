package overlap

import (
    "testing"
    "time"
)

// mustParse is a helper to parse RFC3339 timestamps in tests.
func mustParse(value string) time.Time {
    t, err := time.Parse(time.RFC3339, value)
    if err != nil {
        panic(err)
    }
    return t
}

func TestOverlaps(t *testing.T) {
    tests := []struct {
        name        string
        aStart      string
        aEnd        string
        bStart      string
        bEnd        string
        wantOverlap bool
    }{
        {
            name:        "no overlap, a before b",
            aStart:      "2024-01-01T00:00:00Z",
            aEnd:        "2024-01-10T00:00:00Z",
            bStart:      "2024-01-11T00:00:00Z",
            bEnd:        "2024-01-20T00:00:00Z",
            wantOverlap: false,
        },
        {
            name:        "no overlap, b before a",
            aStart:      "2024-01-11T00:00:00Z",
            aEnd:        "2024-01-20T00:00:00Z",
            bStart:      "2024-01-01T00:00:00Z",
            bEnd:        "2024-01-10T00:00:00Z",
            wantOverlap: false,
        },
        {
            name:        "overlap case",
            aStart:      "2024-01-05T00:00:00Z",
            aEnd:        "2024-01-15T00:00:00Z",
            bStart:      "2024-01-10T00:00:00Z",
            bEnd:        "2024-01-20T00:00:00Z",
            wantOverlap: true,
        },
        {
            name:        "touching boundaries",
            aStart:      "2024-01-01T00:00:00Z",
            aEnd:        "2024-01-05T00:00:00Z",
            bStart:      "2024-01-05T00:00:00Z",
            bEnd:        "2024-01-10T00:00:00Z",
            wantOverlap: true,
        },
        {
            name:        "identical ranges",
            aStart:      "2024-01-01T00:00:00Z",
            aEnd:        "2024-01-01T00:00:00Z",
            bStart:      "2024-01-01T00:00:00Z",
            bEnd:        "2024-01-01T00:00:00Z",
            wantOverlap: true,
        },
    }

    for _, tc := range tests {
        tc := tc // capture range variable
        t.Run(tc.name, func(t *testing.T) {
            a := DateRange{Start: mustParse(tc.aStart), End: mustParse(tc.aEnd)}
            b := DateRange{Start: mustParse(tc.bStart), End: mustParse(tc.bEnd)}
            got := Overlaps(a, b)
            if got != tc.wantOverlap {
                t.Errorf("%s: Overlaps(%v-%v, %v-%v) = %v; want %v",
                    tc.name, tc.aStart, tc.aEnd, tc.bStart, tc.bEnd, got, tc.wantOverlap)
            }
        })
    }
} 