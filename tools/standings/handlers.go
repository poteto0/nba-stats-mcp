package standings

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/poteto0/nba-stats-mcp/internal"
)

func GetStandings(ctx context.Context, req *mcp.CallToolRequest, input GetStandingsInput) (
	*mcp.CallToolResult, GetStandingsToolsResult, error,
) {
	if err := ValidateGetStandingsInput(&input); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Invalid input: " + err.Error()},
			},
		}, GetStandingsToolsResult{}, nil
	}

	client := internal.GetGNSClient()
	result := client.Stats.GetLeagueStandings(
		&types.LeagueStandingsParams{
			LeagueID:   "00",
			Season:     input.Season,
			SeasonType: "Regular Season",
		},
	)
	if result.Error != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch league standings."},
			},
		}, GetStandingsToolsResult{}, nil
	}

	if input.Conference == "" {
		return nil, GetStandingsToolsResult{Standings: result.Contents.Standings[:input.Limit]}, nil
	}

	var filteredStandings []types.LeagueStandingsRecord
	for _, record := range result.Contents.Standings {
		if record.Conference == input.Conference {
			filteredStandings = append(filteredStandings, record)
		}
	}

	return nil, GetStandingsToolsResult{Standings: filteredStandings[:input.Limit]}, nil
}

func GetStandingsSummary(ctx context.Context, req *mcp.CallToolRequest, input GetStandingsInput) (
	*mcp.CallToolResult, GetStandingsSummaryResult, error,
) {
	if err := ValidateGetStandingsInput(&input); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Invalid input: " + err.Error()},
			},
		}, GetStandingsSummaryResult{}, nil
	}

	client := internal.GetGNSClient()
	result := client.Stats.GetLeagueStandings(
		&types.LeagueStandingsParams{
			LeagueID:   "00",
			Season:     input.Season,
			SeasonType: "Regular Season",
		},
	)
	if result.Error != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch league standings."},
			},
		}, GetStandingsSummaryResult{}, nil
	}

	summary := make([]GetStandingsSummaryRecord, input.Limit)
	cnt := 0
	for _, record := range result.Contents.Standings {
		if cnt >= input.Limit {
			break
		}

		if input.Conference != "" && record.Conference != input.Conference {
			continue
		}

		summary[cnt] = GetStandingsSummaryRecord{
			TeamName:   record.TeamName,
			Conference: record.Conference,
			Wins:       record.Wins,
			Losses:     record.Losses,
		}
		cnt++
	}

	return nil, GetStandingsSummaryResult{Standings: summary}, nil
}

func GetClutchStandings(ctx context.Context, req *mcp.CallToolRequest, input GetStandingsInput) (
	*mcp.CallToolResult, ClutchStandingsResult, error,
) {
	if err := ValidateGetStandingsInput(&input); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Invalid input: " + err.Error()},
			},
		}, ClutchStandingsResult{}, nil
	}

	client := internal.GetGNSClient()
	result := client.Stats.GetLeagueStandings(
		&types.LeagueStandingsParams{
			LeagueID:   "00",
			Season:     input.Season,
			SeasonType: "Regular Season",
		},
	)
	if result.Error != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch league standings."},
			},
		}, ClutchStandingsResult{}, nil
	}

	clutchRecords := make([]ClutchStandingsRecord, input.Limit)
	cnt := 0
	for _, record := range result.Contents.Standings {
		if cnt >= input.Limit {
			break
		}

		if input.Conference != "" && record.Conference != input.Conference {
			continue
		}

		clutchRecords[cnt] = ClutchStandingsRecord{
			TeamName:       record.TeamName,
			Conference:     record.Conference,
			ThreePTSOrLess: record.ThreePTSOrLess,
			OT:             record.OT,
			AheadAtHalf:    record.AheadAtHalf,
			BehindAtHalf:   record.BehindAtHalf,
			AheadAtThird:   record.AheadAtThird,
			BehindAtThird:  record.BehindAtThird,
		}
		cnt++
	}

	return nil, ClutchStandingsResult{Standings: clutchRecords}, nil
}

func GetMonthlyStandings(ctx context.Context, req *mcp.CallToolRequest, input GetStandingsInput) (
	*mcp.CallToolResult, MonthlyStandingsResult, error,
) {
	if err := ValidateGetStandingsInput(&input); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Invalid input: " + err.Error()},
			},
		}, MonthlyStandingsResult{}, nil
	}

	client := internal.GetGNSClient()
	result := client.Stats.GetLeagueStandings(
		&types.LeagueStandingsParams{
			LeagueID:   "00",
			Season:     input.Season,
			SeasonType: "Regular Season",
		},
	)
	if result.Error != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch league standings."},
			},
		}, MonthlyStandingsResult{}, nil
	}

	monthlyRecords := make([]MonthlyStandingsRecord, input.Limit)
	cnt := 0
	for _, record := range result.Contents.Standings {
		if cnt >= input.Limit {
			break
		}

		if input.Conference != "" && record.Conference != input.Conference {
			continue
		}

		monthlyEntries := []MonthlyRecordEntry{}
		if record.Oct != nil {
			monthlyEntries = append(monthlyEntries, MonthlyRecordEntry{Month: "oct", Record: *record.Oct})
		}
		if record.Nov != nil {
			monthlyEntries = append(monthlyEntries, MonthlyRecordEntry{Month: "nov", Record: *record.Nov})
		}
		if record.Dec != nil {
			monthlyEntries = append(monthlyEntries, MonthlyRecordEntry{Month: "dec", Record: *record.Dec})
		}
		if record.Jan != nil {
			monthlyEntries = append(monthlyEntries, MonthlyRecordEntry{Month: "jan", Record: *record.Jan})
		}
		if record.Feb != nil {
			monthlyEntries = append(monthlyEntries, MonthlyRecordEntry{Month: "feb", Record: *record.Feb})
		}
		if record.Mar != nil {
			monthlyEntries = append(monthlyEntries, MonthlyRecordEntry{Month: "mar", Record: *record.Mar})
		}
		if record.Apr != nil {
			monthlyEntries = append(monthlyEntries, MonthlyRecordEntry{Month: "apr", Record: *record.Apr})
		}

		monthlyRecords[cnt] = MonthlyStandingsRecord{
			TeamName:   record.TeamName,
			Conference: record.Conference,
			Monthly:    monthlyEntries,
		}
		cnt++
	}

	return nil, MonthlyStandingsResult{Standings: monthlyRecords}, nil
}
