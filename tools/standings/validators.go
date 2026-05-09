package standings

import "fmt"

func ValidateGetStandingsInput(input *GetStandingsInput) error {
	if len(input.Season) != 0 && (len(input.Season) != 7 || input.Season[4] != '-') {
		return fmt.Errorf("invalid season format: %s", input.Season)
	}

	if input.Limit == 0 {
		input.Limit = 30
	}

	if input.Limit < 0 || input.Limit > 30 {
		return fmt.Errorf("limit must be between 0 and 30, got %d", input.Limit)
	}

	if len(input.Conference) != 0 && input.Conference != "East" && input.Conference != "West" {
		return fmt.Errorf("conference must be 'East', 'West', or empty, got '%s'", input.Conference)
	}

	if input.Conference != "" && input.Limit > 15 {
		return fmt.Errorf("limit cannot exceed 15 when filtering by conference, got %d", input.Limit)
	}

	return nil
}
