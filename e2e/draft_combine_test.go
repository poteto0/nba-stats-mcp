package e2e

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/poteto0/nba-stats-mcp/tools/draft"
	"github.com/stretchr/testify/assert"
)

func TestGetDraftCombineHappyPath(t *testing.T) {
	// Arrange
	req := &mcp.CallToolRequest{}
	ctx := context.Background()
	input := draft.GetDraftCombineInput{
		SeasonYear: "2023-24",
	}

	// Act
	_, result, err := draft.GetDraftCombine(ctx, req, input)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, result.CombineStats)
}

func TestGetCombineSimilarityHappyPath(t *testing.T) {
	// Arrange
	req := &mcp.CallToolRequest{}
	ctx := context.Background()
	input := draft.GetCombineSimilarityInput{
		PlayerSeasonYear:   "2024-25",
		PlayerName:         "Donovan Clingan",
		TopK:               5,
		SeasonYearMoreThan: "2019-20",
	}

	// Act
	_, result, err := draft.GetCombineSimilarity(ctx, req, input)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, result.SimilarPlayers)
	assert.LessOrEqual(t, len(result.SimilarPlayers), 5)
}
