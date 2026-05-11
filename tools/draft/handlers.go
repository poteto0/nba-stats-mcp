package draft

import (
	"context"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
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

	if emptyQuery(input) {
		return nil, GetDraftCombineStatsResult{CombineStats: result.Contents.CombineStats}, nil
	}

	var filteredStats []types.DraftCombineStatsRecord
	for _, record := range result.Contents.CombineStats {
		if input.PlayerName != "" && !strings.Contains(strings.ToLower(record.PlayerName), strings.ToLower(input.PlayerName)) {
			continue
		}
		if len(input.Positions) > 0 {
			matchesPosition := false
			for _, pos := range input.Positions {
				if strings.EqualFold(record.Position, pos) {
					matchesPosition = true
					break
				}
			}
			if !matchesPosition {
				continue
			}
		}
		if input.MoreThanHeightWShoesInches > 0 && (record.HeightWShoes == nil || *record.HeightWShoes <= input.MoreThanHeightWShoesInches) {
			continue
		}
		if input.MoreThanWeightPounds > 0 && (record.Weight == nil || *record.Weight <= input.MoreThanWeightPounds) {
			continue
		}
		if input.MoreThanWingspanInches > 0 && (record.Wingspan == nil || *record.Wingspan <= input.MoreThanWingspanInches) {
			continue
		}
		if input.MoreThanVerticalStandingInches > 0 && (record.StandingVertical == nil || *record.StandingVertical <= input.MoreThanVerticalStandingInches) {
			continue
		}
		if input.MoreThanMaxVerticalInches > 0 && (record.MaxVertical == nil || *record.MaxVertical <= input.MoreThanMaxVerticalInches) {
			continue
		}
		if input.MoreThanStandingReachInches > 0 && (record.StandingReach == nil || *record.StandingReach <= input.MoreThanStandingReachInches) {
			continue
		}
		if input.FasterThanLaneAgilitySeconds > 0 && (record.LaneAgility == nil || *record.LaneAgility >= input.FasterThanLaneAgilitySeconds) {
			continue
		}
		filteredStats = append(filteredStats, record)
	}

	return nil, GetDraftCombineStatsResult{CombineStats: filteredStats}, nil
}

func emptyQuery(input GetDraftCombineInput) bool {
	return input.PlayerName == "" &&
		len(input.Positions) == 0 &&
		input.MoreThanHeightWShoesInches == 0 &&
		input.MoreThanWeightPounds == 0 &&
		input.MoreThanWingspanInches == 0 &&
		input.MoreThanVerticalStandingInches == 0 &&
		input.MoreThanMaxVerticalInches == 0 &&
		input.MoreThanStandingReachInches == 0 &&
		input.FasterThanLaneAgilitySeconds == 0
}
