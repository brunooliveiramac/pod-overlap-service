package dto

import "testing"

func TestDateRangeDTOValidate(t *testing.T) {
    tests := []struct {
        name    string
        dto     DateRangeDTO
        wantErr bool
    }{
        {"valid range", DateRangeDTO{Start: "2024-01-01T00:00:00Z", End: "2024-01-05T00:00:00Z"}, false},
        {"equal range", DateRangeDTO{Start: "2024-01-01T00:00:00Z", End: "2024-01-01T00:00:00Z"}, false},
        {"start after end", DateRangeDTO{Start: "2024-01-05T00:00:00Z", End: "2024-01-01T00:00:00Z"}, true},
        {"invalid start format", DateRangeDTO{Start: "not-a-date", End: "2024-01-01T00:00:00Z"}, true},
        {"invalid end format", DateRangeDTO{Start: "2024-01-01T00:00:00Z", End: "bad-date"}, true},
    }
    for _, tc := range tests {
        tc := tc
        t.Run(tc.name, func(t *testing.T) {
            err := tc.dto.Validate()
            if (err != nil) != tc.wantErr {
                t.Fatalf("%s: Validate() error = %v, wantErr %v", tc.name, err, tc.wantErr)
            }
        })
    }
}

func TestOverlapRequestDTOValidate(t *testing.T) {
    valid := OverlapRequestDTO{
        StartRange: DateRangeDTO{Start: "2024-01-01T00:00:00Z", End: "2024-01-02T00:00:00Z"},
        EndRange:   DateRangeDTO{Start: "2024-01-03T00:00:00Z", End: "2024-01-04T00:00:00Z"},
    }
    if err := valid.Validate(); err != nil {
        t.Errorf("Valid request Validate() returned error: %v", err)
    }

    tests := []struct {
        name    string
        req     OverlapRequestDTO
        wantErr bool
    }{
        {"invalid start range", OverlapRequestDTO{StartRange: DateRangeDTO{Start: "2024-01-05T00:00:00Z", End: "2024-01-04T00:00:00Z"}, EndRange: valid.EndRange}, true},
        {"invalid end range", OverlapRequestDTO{StartRange: valid.StartRange, EndRange: DateRangeDTO{Start: "2024-01-05T00:00:00Z", End: "2024-01-04T00:00:00Z"}}, true},
    }
    for _, tc := range tests {
        tc := tc
        t.Run(tc.name, func(t *testing.T) {
            err := tc.req.Validate()
            if (err != nil) != tc.wantErr {
                t.Fatalf("%s: Validate() error = %v, wantErr %v", tc.name, err, tc.wantErr)
            }
        })
    }
} 