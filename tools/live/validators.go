package live

import "fmt"

func ValidateBoxScoreInput(input *BoxScoreInput) error {
	if input == nil {
		return fmt.Errorf("input cannot be nil")
	}

	if input.GameId == "" {
		return fmt.Errorf("game ID is required")
	}

	if len(input.GameId) != 10 {
		return fmt.Errorf("game ID must be 10 characters long")
	}

	return nil
}

func ValidatePlayByPlayInput(input *PlayByPlayInput) error {
	if input == nil {
		return fmt.Errorf("input cannot be nil")
	}

	if input.GameId == "" {
		return fmt.Errorf("game ID is required")
	}

	if len(input.GameId) != 10 {
		return fmt.Errorf("game ID must be 10 characters long")
	}

	return nil
}

func ValidateGetSpecificPlayByPlayDetailsInput(input *GetSpecificPlayByPlayDetailsInput) error {
	if input == nil {
		return fmt.Errorf("input cannot be nil")
	}

	if input.GameId == "" {
		return fmt.Errorf("game ID is required")
	}

	if len(input.GameId) != 10 {
		return fmt.Errorf("game ID must be 10 characters long")
	}

	if len(input.ActionNumbers) == 0 {
		return fmt.Errorf("at least one action number must be provided")
	}

	return nil
}
