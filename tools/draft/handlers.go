package draft

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/poteto-go/tslice"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/poteto0/nba-stats-mcp/internal"
)

func GetDraftCombine(ctx context.Context, req *mcp.CallToolRequest, input GetDraftCombineInput) (
	*mcp.CallToolResult, GetDraftCombineStatsResult, error,
) {
	if err := ValidateGetDraftCombineInput(&input); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Invalid input: " + err.Error()},
			},
		}, GetDraftCombineStatsResult{}, nil
	}

	client := internal.GetGNSClient()
	result := client.Draft.GetCombineStats(
		&types.DraftCombineStatsParams{
			SeasonYear: input.SeasonYear,
		},
	)

	if result.Error != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch draft combine results."},
			},
		}, GetDraftCombineStatsResult{}, nil
	}

	filteredStats := filterDraftCombineStatsRecords(result.Contents.CombineStats, input)

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "Draft combine results retrieved successfully."},
		},
	}, GetDraftCombineStatsResult{CombineStats: filteredStats}, nil
}

func GetCombineSimilarity(ctx context.Context, req *mcp.CallToolRequest, input GetCombineSimilarityInput) (
	*mcp.CallToolResult, GetCombineSimilarityResult, error,
) {
	if err := ValidateGetCombineSimilarityInput(&input); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Invalid input: " + err.Error()},
			},
		}, GetCombineSimilarityResult{}, nil
	}

	client := internal.GetGNSClient()
	result := client.Draft.GetCombineStats(
		&types.DraftCombineStatsParams{
			SeasonYear: input.PlayerSeasonYear,
		},
	)

	if result.Error != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch draft combine results."},
			},
		}, GetCombineSimilarityResult{}, nil
	}

	filteredStats := filterDraftCombineStatsRecords(
		result.Contents.CombineStats,
		GetDraftCombineInput{
			PlayerName: input.PlayerName,
		},
	)
	if len(filteredStats) == 0 {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "No draft combine results found for the specified player and season."},
			},
		}, GetCombineSimilarityResult{}, nil
	}

	playerStats := filteredStats[0]

	seasonYearsToSearch, err := generateSeasonYearRange(
		input.PlayerSeasonYear,
		input.SeasonYearMoreThan,
		input.SeasonLessThan,
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Invalid season year range: " + err.Error()},
			},
		}, GetCombineSimilarityResult{}, nil
	}

	var similarityResults []DraftCombineSimilarityRecord
	for _, seasonYear := range seasonYearsToSearch {
		result := client.Draft.GetCombineStats(
			&types.DraftCombineStatsParams{
				SeasonYear: seasonYear,
			},
		)

		if result.Error != nil {
			continue
		}

		records := result.Contents.CombineStats
		for _, record := range records {
			similarityScore := calculateSimilarityScore(playerStats, record)
			similarityResults = append(similarityResults, DraftCombineSimilarityRecord{
				DraftCombineStatsRecord: record,
				SimilarityScore:         similarityScore,
			})
		}
	}

	tslice.Sort(similarityResults, func(a, b DraftCombineSimilarityRecord) int {
		if a.SimilarityScore < b.SimilarityScore {
			return -1
		} else if a.SimilarityScore > b.SimilarityScore {
			return 1
		}
		return 0
	})

	if len(similarityResults) > input.TopK {
		similarityResults = similarityResults[:input.TopK]
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "Similar players retrieved successfully."},
		},
	}, GetCombineSimilarityResult{SimilarPlayers: similarityResults}, nil
}
