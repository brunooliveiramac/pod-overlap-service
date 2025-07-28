package httpapi_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	httpapi "github.com/brunooliveiramac/pod-overlap-service/internal/platform/http"
	dto "github.com/brunooliveiramac/pod-overlap-service/internal/platform/http/dto"
	"github.com/gin-gonic/gin"
)

func TestE2EOverlap(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	httpapi.RegisterRoutes(router)

	server := httptest.NewServer(router)
	defer server.Close()

	type Resp struct {
		Overlap bool `json:"overlap"`
	}

	tests := []struct {
		name        string
		payload     dto.OverlapRequestDTO
		wantStatus  int
		wantOverlap bool
	}{
		{
			name:        "valid overlap",
			payload:     dto.OverlapRequestDTO{StartRange: dto.DateRangeDTO{Start: "2024-07-01T00:00:00Z", End: "2024-07-10T00:00:00Z"}, EndRange: dto.DateRangeDTO{Start: "2024-07-05T00:00:00Z", End: "2024-07-15T00:00:00Z"}},
			wantStatus:  http.StatusOK,
			wantOverlap: true,
		},
		{
			name:       "invalid start end order",
			payload:    dto.OverlapRequestDTO{StartRange: dto.DateRangeDTO{Start: "2024-07-10T00:00:00Z", End: "2024-07-05T00:00:00Z"}, EndRange: dto.DateRangeDTO{Start: "2024-07-05T00:00:00Z", End: "2024-07-15T00:00:00Z"}},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes, err := json.Marshal(tc.payload)
			if err != nil {
				t.Fatalf("json.Marshal error: %v", err)
			}
			resp, err := http.Post(server.URL+"/api/overlap", "application/json", bytes.NewReader(bodyBytes))
			if err != nil {
				t.Fatalf("HTTP POST error: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.wantStatus {
				t.Errorf("expected status %d, got %d", tc.wantStatus, resp.StatusCode)
			}
			if resp.StatusCode == http.StatusOK {
				var r Resp
				if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
					t.Fatalf("decoding response error: %v", err)
				}
				if r.Overlap != tc.wantOverlap {
					t.Errorf("expected overlap %v, got %v", tc.wantOverlap, r.Overlap)
				}
			}
		})
	}
}
