package dto

import (
	"fmt"
	"time"
)

type OverlapRequestDTO struct {
	StartRange DateRangeDTO `json:"start_range" binding:"required"`
	EndRange   DateRangeDTO `json:"end_range" binding:"required"`
}

func (r *OverlapRequestDTO) Validate() error {
	if err := r.StartRange.Validate(); err != nil {
		return fmt.Errorf("start_range: %w", err)
	}
	if err := r.EndRange.Validate(); err != nil {
		return fmt.Errorf("end_range: %w", err)
	}
	return nil
}

type DateRangeDTO struct {
	Start string `json:"start" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	End   string `json:"end" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
}

// Validate checks that Start <= End.
func (d *DateRangeDTO) Validate() error {
	start, err := time.Parse(time.RFC3339, d.Start)
	if err != nil {
		return fmt.Errorf("invalid start date: %w", err)
	}
	end, err := time.Parse(time.RFC3339, d.End)
	if err != nil {
		return fmt.Errorf("invalid end date: %w", err)
	}
	if start.After(end) {
		return fmt.Errorf("start %s must be before or equal to end %s", d.Start, d.End)
	}
	return nil
}

type OverlapResponseDTO struct {
	Overlap bool `json:"overlap"`
}
