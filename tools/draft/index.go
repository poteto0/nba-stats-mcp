package draft

import "github.com/modelcontextprotocol/go-sdk/mcp"

func RegisterTools(server *mcp.Server) *mcp.Server {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_nba_draft_combine",
		Description: "Returns the NBA Draft Combine results based on the provided input parameters. The input can include filters such as height with shoes, weight in pounds, wingspan inches, vertical standing inches, max vertical inches, and lane agility seconds. The output will include a list of players who meet the specified criteria along with their combine results.",
	}, GetDraftCombine)

	return server
}
