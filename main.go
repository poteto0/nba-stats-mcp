package main

import (
	"log"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/poteto0/nba-stats-mcp/tools"
)

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

	http.Handle("/", handler)

	log.Printf("MCP server listening on :%s/", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
