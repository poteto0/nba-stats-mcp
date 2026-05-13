package draft

import (
	"reflect"
	"testing"
)

func TestGenerateSeasonYearRange(t *testing.T) {
	tests := []struct {
		name               string
		playerSeasonYear   string
		seasonYearMoreThan string
		seasonYearLessThan string
		want               []string
		wantErr            bool
	}{
		{
			name:               "default search range is one year",
			playerSeasonYear:   "2023-24",
			seasonYearMoreThan: "",
			seasonYearLessThan: "",
			want:               []string{"2022-23", "2024-25"},
			wantErr:            false,
		},
		{
			name:               "SeasonYear is minimum (2000-01)",
			playerSeasonYear:   "2000-01",
			seasonYearMoreThan: "",
			seasonYearLessThan: "",
			want:               []string{"2001-02"},
			wantErr:            false,
		},
		{
			name:               "specified range",
			playerSeasonYear:   "2023-24",
			seasonYearMoreThan: "2020-21",
			seasonYearLessThan: "2024-25",
			want:               []string{"2020-21", "2021-22", "2022-23", "2024-25"},
			wantErr:            false,
		},
		{
			name:               "error: start after end",
			playerSeasonYear:   "2023-24",
			seasonYearMoreThan: "2024-25",
			seasonYearLessThan: "",
			wantErr:            true,
		},
		{
			name:               "error: end before start",
			playerSeasonYear:   "2023-24",
			seasonYearMoreThan: "",
			seasonYearLessThan: "2022-23",
			wantErr:            true,
		},
		{
			name:               "error: before minimum year",
			playerSeasonYear:   "2023-24",
			seasonYearMoreThan: "1999-00",
			seasonYearLessThan: "",
			wantErr:            true,
		},
		{
			name:               "error: after maximum year",
			playerSeasonYear:   "2023-24",
			seasonYearMoreThan: "",
			seasonYearLessThan: "2027-28",
			wantErr:            true,
		},
		{
			name:               "seasonYear must be between 2000-01 and 2026-27",
			playerSeasonYear:   "1999-00",
			seasonYearMoreThan: "",
			seasonYearLessThan: "",
			wantErr:            true,
		},
		{
			name:               "seasonYear must be between 2000-01 and 2026-27",
			playerSeasonYear:   "2027-28",
			seasonYearMoreThan: "",
			seasonYearLessThan: "",
			wantErr:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateSeasonYearRange(tt.playerSeasonYear, tt.seasonYearMoreThan, tt.seasonYearLessThan)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateSeasonYearRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateSeasonYearRange() got = %v, want %v", got, tt.want)
			}
		})
	}
}
