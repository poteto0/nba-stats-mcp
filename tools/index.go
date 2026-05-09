package tools

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/poteto0/nba-stats-mcp/tools/standings"
)

func RegisterTools(server *mcp.Server) *mcp.Server {
	standings.RegisterTools(server)
	return server
}
