package draft

import "fmt"

func ValidateGetDraftCombineInput(input *GetDraftCombineInput) error {
	if input == nil {
		return fmt.Errorf("input cannot be nil")
	}

	if input.MoreThanHeightWShoesInches < 0 {
		return fmt.Errorf("height with shoes inches cannot be negative")
	}

	if input.MoreThanWeightPounds < 0 {
		return fmt.Errorf("weight in pounds cannot be negative")
	}

	if input.MoreThanWingspanInches < 0 {
		return fmt.Errorf("wingspan inches cannot be negative")
	}

	if input.MoreThanVerticalStandingInches < 0 {
		return fmt.Errorf("vertical standing inches cannot be negative")
	}

	if input.MoreThanMaxVerticalInches < 0 {
		return fmt.Errorf("max vertical inches cannot be negative")
	}

	if input.FasterThanLaneAgilitySeconds < 0 {
		return fmt.Errorf("lane agility seconds cannot be negative")
	}

	return nil
}
