package middleware

import (
	"net/http"
	"time"

	"github.com/poteto0/nba-stats-mcp/constants"
	"golang.org/x/time/rate"
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(rate.Every(time.Minute/time.Duration(constants.MaxRequestsPerMinute)), constants.MaxRequestsPerMinute)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
