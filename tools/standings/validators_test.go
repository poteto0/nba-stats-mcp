package standings_test

import (
	"testing"

	"github.com/poteto0/nba-stats-mcp/tools/standings"
)

func TestValidateGetStandingsInput(t *testing.T) {
	tests := []struct {
		name    string
		input   standings.GetStandingsInput
		wantErr bool
	}{
		{
			name:    "no input provided",
			input:   standings.GetStandingsInput{},
			wantErr: false,
		},
		{
			name: "valid season format",
			input: standings.GetStandingsInput{
				Season:     "2025-26",
				Limit:      30,
				Conference: "East",
			},
			wantErr: false,
		},
		{
			name:    "invalid season format - missing dash",
			input:   standings.GetStandingsInput{Season: "202526", Conference: ""},
			wantErr: true,
		},
		{
			name:    "invalid season format - wrong length",
			input:   standings.GetStandingsInput{Season: "2025-2026", Limit: 30},
			wantErr: true,
		},
		{
			name:    "invalid limit - negative",
			input:   standings.GetStandingsInput{Season: "2025-26", Limit: -1},
			wantErr: true,
		},
		{
			name:    "invalid limit - exceeds maximum",
			input:   standings.GetStandingsInput{Season: "2025-26", Limit: 31},
			wantErr: true,
		},
		{
			name:    "invalid conference - not East or West",
			input:   standings.GetStandingsInput{Season: "2025-26", Conference: "North"},
			wantErr: true,
		},
		{
			name:    "invalid limit when conference filter is applied",
			input:   standings.GetStandingsInput{Season: "2025-26", Limit: 20, Conference: "East"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := standings.ValidateGetStandingsInput(&tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateGetStandingsInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
