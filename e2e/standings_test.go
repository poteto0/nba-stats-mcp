package e2e

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/poteto0/nba-stats-mcp/tools/standings"
	"github.com/stretchr/testify/assert"
)

func TestGetStandingsHappyPath(t *testing.T) {
	// Arrange
	req := &mcp.CallToolRequest{}
	ctx := context.Background()
	input := standings.GetStandingsInput{Season: "2023-24", Limit: 5}

	// Act
	_, _, err := standings.GetStandings(ctx, req, input)

	// Assert
	assert.NoError(t, err)
}

func TestGetStandingsSummaryHappyPath(t *testing.T) {
	// Arrange
	req := &mcp.CallToolRequest{}
	ctx := context.Background()
	input := standings.GetStandingsInput{Season: "2023-24", Limit: 5}

	// Act
	_, _, err := standings.GetStandingsSummary(ctx, req, input)

	// Assert
	assert.NoError(t, err)
}

func TestGetClutchStandingsHappyPath(t *testing.T) {
	// Arrange
	req := &mcp.CallToolRequest{}
	ctx := context.Background()
	input := standings.GetStandingsInput{Season: "2023-24", Limit: 5}

	// Act
	_, _, err := standings.GetClutchStandings(ctx, req, input)

	// Assert
	assert.NoError(t, err)
}

func TestGetMonthlyStandingsHappyPath(t *testing.T) {
	// Arrange
	req := &mcp.CallToolRequest{}
	ctx := context.Background()
	input := standings.GetStandingsInput{Season: "2023-24", Limit: 5}

	// Act
	_, _, err := standings.GetMonthlyStandings(ctx, req, input)

	// Assert
	assert.NoError(t, err)
}
