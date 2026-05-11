package tools

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/poteto0/nba-stats-mcp/tools/draft"
	"github.com/poteto0/nba-stats-mcp/tools/live"
	"github.com/poteto0/nba-stats-mcp/tools/standings"
)

func RegisterTools(server *mcp.Server) *mcp.Server {
	standings.RegisterTools(server)
	live.RegisterTools(server)
	draft.RegisterTools(server)
	return server
}
