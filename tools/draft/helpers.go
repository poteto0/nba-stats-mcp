package draft

import (
	"errors"
	"strings"

	"github.com/poteto0/go-nba-sdk/types"
	"github.com/poteto0/nba-stats-mcp/internal"
)

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

func filterDraftCombineStatsRecords(records []types.DraftCombineStatsRecord, input GetDraftCombineInput) []types.DraftCombineStatsRecord {
	if emptyQuery(input) {
		return records
	}

	var filteredStats []types.DraftCombineStatsRecord
	for _, record := range records {
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

	return filteredStats
}

// Max = 2026-27, Min = 2001-02
func generateSeasonYearRange(
	playerSeasonYear string,
	seasonYearMoreThan string,
	seasonYearLessThan string,
) ([]string, error) {
	const (
		minYear = 2000
		maxYear = 2026
	)

	playerSeasonBeginYear, err := internal.ParseSeasonBeginYear(playerSeasonYear)
	if err != nil {
		return nil, err
	}

	if playerSeasonBeginYear < minYear || playerSeasonBeginYear > maxYear {
		return nil, errors.New("playerSeasonYear must be between 2000-01 and 2026-27")
	}

	leftSeasonYears, err := underSeasonRange(playerSeasonBeginYear, seasonYearMoreThan)
	if err != nil {
		return nil, err
	}

	rightSeasonYears, err := overSeasonRange(playerSeasonBeginYear, seasonYearLessThan)
	if err != nil {
		return nil, err
	}

	return append(leftSeasonYears, rightSeasonYears...), nil
}

// SeasonYearMoreThan <= x < SeasonYear
func underSeasonRange(seasonBeginYear int, seasonYearMoreThan string) ([]string, error) {
	const minBeginYear = 2000
	if seasonBeginYear <= minBeginYear {
		return []string{}, nil
	}

	if seasonYearMoreThan == "" {
		seasonYearMoreThan = internal.FormatSeasonYear(seasonBeginYear - 1)
	}
	seasonYearMoreThanBeginYear, err := internal.ParseSeasonBeginYear(seasonYearMoreThan)
	if err != nil {
		return nil, err
	}

	if seasonYearMoreThanBeginYear < minBeginYear {
		return nil, errors.New("seasonYearMoreThan must be less than minimum season year (2000-01)")
	}

	if seasonYearMoreThanBeginYear >= seasonBeginYear {
		return nil, errors.New("seasonYearMoreThan must be less than playerSeasonYear")
	}

	var seasonYears []string
	for i := seasonYearMoreThanBeginYear; i < seasonBeginYear; i++ {
		seasonYears = append(seasonYears, internal.FormatSeasonYear(i))
	}

	return seasonYears, nil
}

func overSeasonRange(seasonBeginYear int, seasonYearLessThan string) ([]string, error) {
	const maxBeginYear = 2026
	if seasonBeginYear >= maxBeginYear {
		return []string{}, nil
	}

	if seasonYearLessThan == "" {
		seasonYearLessThan = internal.FormatSeasonYear(seasonBeginYear + 1)
	}
	seasonYearLessThanBeginYear, err := internal.ParseSeasonBeginYear(seasonYearLessThan)
	if err != nil {
		return nil, err
	}

	if seasonYearLessThanBeginYear > maxBeginYear {
		return nil, errors.New("seasonYearLessThan must be less than maximum season year (2026-27)")
	}

	if seasonYearLessThanBeginYear <= seasonBeginYear {
		return nil, errors.New("seasonYearLessThan must be greater than playerSeasonYear")
	}

	var seasonYears []string
	for i := seasonBeginYear + 1; i <= seasonYearLessThanBeginYear; i++ {
		seasonYears = append(seasonYears, internal.FormatSeasonYear(i))
	}

	return seasonYears, nil
}
