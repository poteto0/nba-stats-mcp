package live_test

import (
	"testing"

	"github.com/poteto0/nba-stats-mcp/tools/live"
)

func TestValidateBoxScoreInput(t *testing.T) {
	tests := []struct {
		name    string
		input   *live.BoxScoreInput
		wantErr bool
	}{
		{
			name:    "nil input is error",
			input:   nil,
			wantErr: true,
		},
		{
			name:    "no game ID provided",
			input:   &live.BoxScoreInput{},
			wantErr: true,
		},
		{
			name:    "game ID too short",
			input:   &live.BoxScoreInput{GameId: "12345"},
			wantErr: true,
		},
		{
			name:    "game ID too long",
			input:   &live.BoxScoreInput{GameId: "12345678901"},
			wantErr: true,
		},
		{
			name:    "valid game ID",
			input:   &live.BoxScoreInput{GameId: "1234567890"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := live.ValidateBoxScoreInput(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateBoxScoreInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePlayByPlayInput(t *testing.T) {
	tests := []struct {
		name    string
		input   *live.PlayByPlayInput
		wantErr bool
	}{
		{
			name:    "nil input is error",
			input:   nil,
			wantErr: true,
		},
		{
			name:    "no game ID provided",
			input:   &live.PlayByPlayInput{},
			wantErr: true,
		},
		{
			name:    "game ID too short",
			input:   &live.PlayByPlayInput{GameId: "12345"},
			wantErr: true,
		},
		{
			name:    "game ID too long",
			input:   &live.PlayByPlayInput{GameId: "12345678901"},
			wantErr: true,
		},
		{
			name:    "valid game ID",
			input:   &live.PlayByPlayInput{GameId: "1234567890"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := live.ValidatePlayByPlayInput(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePlayByPlayInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateGetSpecificPlayByPlayDetailsInput(t *testing.T) {
	tests := []struct {
		name    string
		input   *live.GetSpecificPlayByPlayDetailsInput
		wantErr bool
	}{
		{
			name:    "nil input is error",
			input:   nil,
			wantErr: true,
		},
		{
			name:    "no game ID provided",
			input:   &live.GetSpecificPlayByPlayDetailsInput{},
			wantErr: true,
		},
		{
			name:    "game ID too short",
			input:   &live.GetSpecificPlayByPlayDetailsInput{GameId: "12345"},
			wantErr: true,
		},
		{
			name:    "game ID too long",
			input:   &live.GetSpecificPlayByPlayDetailsInput{GameId: "12345678901"},
			wantErr: true,
		},
		{
			name:    "no action numbers provided",
			input:   &live.GetSpecificPlayByPlayDetailsInput{GameId: "1234567890"},
			wantErr: true,
		},
		{
			name:    "valid input",
			input:   &live.GetSpecificPlayByPlayDetailsInput{GameId: "1234567890", ActionNumbers: []int{1, 2, 3}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := live.ValidateGetSpecificPlayByPlayDetailsInput(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateGetSpecificPlayByPlayDetailsInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
