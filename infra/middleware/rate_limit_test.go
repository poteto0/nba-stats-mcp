package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/poteto0/nba-stats-mcp/constants"
)

func TestRateLimitMiddleware(t *testing.T) {
	// Arrange
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	middleware := RateLimitMiddleware(handler)

	// Act & Assert
	// 制限内（MaxRequestsPerMinute）は通過するはず
	for i := 0; i < constants.MaxRequestsPerMinute; i++ {
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		middleware.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("request %d: expected status 200, got %d", i, w.Code)
		}
	}

	// 制限を超えると Too Many Requests が返るはず
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	middleware.ServeHTTP(w, req)
	if w.Code != http.StatusTooManyRequests {
		t.Errorf("expected status 429, got %d", w.Code)
	}
}
