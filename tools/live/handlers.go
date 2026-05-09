package live

import (
	"context"
	"slices"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/poteto0/nba-stats-mcp/internal"
)

func GetScoreBoard(ctx context.Context, req *mcp.CallToolRequest, input GetScoreBoardInput) (
	*mcp.CallToolResult, GetScoreBoardResult, error,
) {
	client := internal.GetGNSClient()
	result := client.Live.GetScoreBoard(nil)

	if result.Error != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch live scoreboard."},
			},
		}, GetScoreBoardResult{}, nil
	}

	return nil, GetScoreBoardResult{Scoreboard: LiveScoreBoardRecord{Games: result.Contents.Scoreboard.Games}}, nil
}

func GetScoreBoardSummary(ctx context.Context, req *mcp.CallToolRequest, input GetScoreBoardInput) (
	*mcp.CallToolResult, GetScoreBoardSummaryResult, error,
) {
	client := internal.GetGNSClient()
	result := client.Live.GetScoreBoard(nil)

	if result.Error != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch live scoreboard."},
			},
		}, GetScoreBoardSummaryResult{}, nil
	}

	summary := LiveScoreBoardSummaryRecord{}
	for _, game := range result.Contents.Scoreboard.Games {
		summary.Games = append(summary.Games, GameSummaryRecord{
			GameId:        game.GameId,
			GameStatus:    game.GameStatusText,
			HomeTeamName:  game.HomeTeam.TeamName,
			AwayTeamName:  game.AwayTeam.TeamName,
			HomeTeamScore: game.HomeTeam.Score,
			AwayTeamScore: game.AwayTeam.Score,
			GameClock:     game.GameClock,
			Period:        game.Period,
		})
	}

	return nil, GetScoreBoardSummaryResult{Scoreboard: summary}, nil
}

func GetGameLeaders(ctx context.Context, req *mcp.CallToolRequest, input GetScoreBoardInput) (
	*mcp.CallToolResult, GetGameLeadersResult, error,
) {
	client := internal.GetGNSClient()
	result := client.Live.GetScoreBoard(nil)

	if result.Error != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch live scoreboard."},
			},
		}, GetGameLeadersResult{}, nil
	}

	gameLeaders := []GameLeaderRecord{}
	for _, game := range result.Contents.Scoreboard.Games {
		gameLeaders = append(gameLeaders, GameLeaderRecord{
			GameId:        game.GameId,
			GameStatus:    game.GameStatusText,
			HomeTeamName:  game.HomeTeam.TeamName,
			AwayTeamName:  game.AwayTeam.TeamName,
			HomeTeamScore: game.HomeTeam.Score,
			AwayTeamScore: game.AwayTeam.Score,
			GameClock:     game.GameClock,
			Period:        game.Period,
			GameLeaders:   game.GameLeaders,
		})
	}

	return nil, GetGameLeadersResult{Scoreboard: gameLeaders}, nil
}

func GetBoxScore(ctx context.Context, req *mcp.CallToolRequest, input BoxScoreInput) (
	*mcp.CallToolResult, GetBoxScoreResult, error,
) {
	client := internal.GetGNSClient()
	result := client.Live.GetBoxScore(&types.BoxScoreParams{GameID: input.GameId})

	if result.Error != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch box score for game ID: " + input.GameId},
			},
		}, GetBoxScoreResult{}, nil
	}

	return nil, GetBoxScoreResult{BoxScore: BoxScoreRecord{Games: result.Contents.Game}}, nil
}

func GetPlayByPlaySummary(ctx context.Context, req *mcp.CallToolRequest, input PlayByPlayInput) (
	*mcp.CallToolResult, GetPlayByPlaySummaryResult, error,
) {
	client := internal.GetGNSClient()
	result := client.Live.GetPlayByPlay(&types.PlayByPlayParams{GameID: input.GameId})

	if result.Error != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch play-by-play data for game ID: " + input.GameId},
			},
		}, GetPlayByPlaySummaryResult{}, nil
	}

	summary := []PlayByPlaySummaryRecord{}
	for _, action := range result.Contents.Game.Actions {
		if input.TeamId != 0 && action.TeamID != input.TeamId {
			continue
		}

		if input.ActionType != "" && action.ActionType != input.ActionType {
			continue
		}

		if input.ShotResult != "" && action.ShotResult != input.ShotResult {
			continue
		}

		if input.Period != 0 && action.Period != input.Period {
			continue
		}

		if input.IsFieldGoal && !(action.IsFieldGoal == 0) {
			continue
		}

		summary = append(summary, PlayByPlaySummaryRecord{
			ActionNumber: action.ActionNumber,
			ActionType:   action.ActionType,
			Period:       action.Period,
			Description:  action.Description,
		})
	}

	return nil, GetPlayByPlaySummaryResult{PlayByPlay: summary}, nil
}

func GetSpecificPlayByPlayDetails(ctx context.Context, req *mcp.CallToolRequest, input GetSpecificPlayByPlayDetailsInput) (
	*mcp.CallToolResult, PlayByPlayDetailsResult, error,
) {
	client := internal.GetGNSClient()
	result := client.Live.GetPlayByPlay(&types.PlayByPlayParams{GameID: input.GameId})

	if result.Error != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch play-by-play data for game ID: " + input.GameId},
			},
		}, PlayByPlayDetailsResult{}, nil
	}

	details := []types.Action{}
	for _, action := range result.Contents.Game.Actions {
		if !slices.Contains(input.ActionNumbers, action.ActionNumber) {
			continue
		}
		details = append(details, action)
	}

	return nil, PlayByPlayDetailsResult{Actions: details}, nil
}
