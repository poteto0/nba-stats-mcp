package standings

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func RegisterTools(server *mcp.Server) *mcp.Server {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_nba_standings",
		Description: "Returns all records related to the leaderboard. In many cases, consider specialized tools to save context.",
	}, GetStandings)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_nba_standings_summary",
		Description: "Return only the team and the win/loss result.",
	}, GetStandingsSummary)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_nba_clutch_standings",
		Description: "Returns records related to performance in close games, such as win-loss record in games decided by 3 points or less, overtime games, and performance when leading or trailing at halftime and after the 3rd quarter.",
	}, GetClutchStandings)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_nba_monthly_standings",
		Description: "Returns monthly win-loss records for each team, allowing analysis of performance trends throughout the season.",
	}, GetMonthlyStandings)

	return server
}
