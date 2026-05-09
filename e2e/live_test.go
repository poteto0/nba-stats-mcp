package e2e

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/poteto0/nba-stats-mcp/tools/live"
	"github.com/stretchr/testify/assert"
)

func TestGetScoreBoardHappyPath(t *testing.T) {
	// Arrange
	req := &mcp.CallToolRequest{}
	ctx := context.Background()
	input := live.GetScoreBoardInput{}

	// Act
	_, _, err := live.GetScoreBoard(ctx, req, input)

	// Assert
	assert.NoError(t, err)
}

func TestGetScoreBoardSummaryHappyPath(t *testing.T) {
	// Arrange
	req := &mcp.CallToolRequest{}
	ctx := context.Background()
	input := live.GetScoreBoardInput{}

	// Act
	_, _, err := live.GetScoreBoardSummary(ctx, req, input)

	// Assert
	assert.NoError(t, err)
}

func TestGetGameLeadersHappyPath(t *testing.T) {
	// Arrange
	req := &mcp.CallToolRequest{}
	ctx := context.Background()
	input := live.GetScoreBoardInput{}

	// Act
	_, _, err := live.GetGameLeaders(ctx, req, input)

	// Assert
	assert.NoError(t, err)
}

func TestGetBoxScoreHappyPath(t *testing.T) {
	// Arrange
	req := &mcp.CallToolRequest{}
	ctx := context.Background()
	// NOTE: Requires a valid game ID for full success, but testing interface callability
	input := live.BoxScoreInput{GameId: "0022300001"}

	// Act
	_, _, err := live.GetBoxScore(ctx, req, input)

	// Assert
	assert.NoError(t, err)
}

func TestGetPlayByPlaySummaryHappyPath(t *testing.T) {
	// Arrange
	req := &mcp.CallToolRequest{}
	ctx := context.Background()
	input := live.PlayByPlayInput{GameId: "0022300001"}

	// Act
	_, _, err := live.GetPlayByPlaySummary(ctx, req, input)

	// Assert
	assert.NoError(t, err)
}

func TestGetSpecificPlayByPlayDetailsHappyPath(t *testing.T) {
	// Arrange
	req := &mcp.CallToolRequest{}
	ctx := context.Background()
	input := live.GetSpecificPlayByPlayDetailsInput{
		GameId:        "0022300001",
		ActionNumbers: []int{1, 2, 3},
	}

	// Act
	_, _, err := live.GetSpecificPlayByPlayDetails(ctx, req, input)

	// Assert
	assert.NoError(t, err)
}
