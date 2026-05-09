package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/poteto0/nba-stats-mcp/constants"
	"github.com/poteto0/nba-stats-mcp/tools"
	"golang.org/x/time/rate"
)

func rateLimitMiddleware(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(rate.Every(time.Minute/time.Duration(constants.MaxRequestsPerMinute)), constants.MaxRequestsPerMinute)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "nba-stats-mcp",
		Version: "0.1.0",
	}, nil)

	tools.RegisterTools(server)

	// Cloud Run は PORT 環境変数でポートを渡してくる
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// StreamableHTTPHandler で HTTP サーバーとして起動
	handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server {
		return server
	}, nil)

	http.Handle("/", rateLimitMiddleware(handler))

	log.Printf("MCP server listening on :%s/", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
