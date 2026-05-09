package live

import "github.com/modelcontextprotocol/go-sdk/mcp"

func RegisterTools(server *mcp.Server) *mcp.Server {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_nba_live_scoreboard",
		Description: "Returns the current live scoreboard data, including all games in progress with details such as team names, scores, time remaining, and other relevant information. In many cases, consider specialized tools to save context.",
	}, GetScoreBoard)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_nba_live_scoreboard_summary",
		Description: "Returns a summary of the current live scoreboard, including all games in progress with key details such as team names, scores, game status, and time remaining.",
	}, GetScoreBoardSummary)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_nba_live_game_leaders",
		Description: "Returns the current live scoreboard data along with the statistical leaders for each game, including points, rebounds, assists, and other relevant stats for both teams.",
	}, GetGameLeaders)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_nba_box_score",
		Description: "Returns the box score data for a specific game, including detailed statistics for each player and team. The input should include the game ID.",
	}, GetBoxScore)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_nba_play_by_play_summary",
		Description: "Returns a summary of the play-by-play data for a specific game, including key actions and their descriptions.",
	}, GetPlayByPlaySummary)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_nba_specific_play_by_play_details",
		Description: "Returns detailed play-by-play data for specific actions in a game, based on the provided action numbers. The input should include the game ID and a list of action numbers to retrieve details for.",
	}, GetSpecificPlayByPlayDetails)

	return server
}
