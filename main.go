package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/poteto0/nba-stats-mcp/infra/middleware"
	"github.com/poteto0/nba-stats-mcp/tools"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "nba-stats-mcp",
		Version: "0.1.1",
	}, nil)

	tools.RegisterTools(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server {
		return server
	}, nil)

	http.Handle("/", middleware.RateLimitMiddleware(handler))

	httpServer := &http.Server{
		Addr:              ":" + port,
		ReadHeaderTimeout: 3 * time.Second,
	}

	log.Printf("MCP server listening on :%s/", port)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
