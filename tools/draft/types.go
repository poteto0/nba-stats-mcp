package draft

import "github.com/poteto0/go-nba-sdk/types"

type GetDraftCombineInput struct {
	SeasonYear                     string   `json:"seasonYear" jsonschema:"The season year for which to retrieve the draft combine results (e.g. '2023-24')"`
	PlayerName                     string   `json:"playerName,omitempty" jsonschema:"Optional filter to only include combine results for players whose name contains the specified string (case-insensitive)"`
	Positions                      []string `json:"positions,omitempty" jsonschema:"Optional filter to only include combine results for players who play the specified positions (e.g. 'PG', 'SG', 'SF', 'PF', 'C'). If you want to filter by guards, you can include both 'PG' and 'SG' in this list."`
	MoreThanHeightWShoesInches     float64  `json:"moreThanHeightInches,omitempty" jsonschema:"Optional filter to only include players with a height greater than the specified number of inches"`
	MoreThanStandingReachInches    float64  `json:"moreThanStandingReachInches,omitempty" jsonschema:"Optional filter to only include players with a standing reach greater than the specified number of inches"`
	MoreThanWeightPounds           float64  `json:"moreThanWeightPounds,omitempty" jsonschema:"Optional filter to only include players with a weight greater than the specified number of pounds"`
	MoreThanWingspanInches         float64  `json:"moreThanWingspanInches,omitempty" jsonschema:"Optional filter to only include players with a wingspan greater than the specified number of inches"`
	MoreThanMaxVerticalInches      float64  `json:"moreThanMaxVerticalInches,omitempty" jsonschema:"Optional filter to only include players with a max vertical leap greater than the specified number of inches"`
	MoreThanVerticalStandingInches float64  `json:"moreThanVerticalStandingInches,omitempty" jsonschema:"Optional filter to only include players with a vertical standing greater than the specified number of inches"`
	FasterThanLaneAgilitySeconds   float64  `json:"fasterThanLaneAgilitySeconds,omitempty" jsonschema:"Optional filter to only include players with a lane agility time faster than the specified number of seconds"`
}

type GetDraftCombineStatsResult struct {
	CombineStats []types.DraftCombineStatsRecord `json:"combineStats" jsonschema:"List of draft combine stats for the specified season, including player name, position, team, and various combine measurements such as height, weight, wingspan, vertical leap, etc."`
}

type GetCombineSimilarityInput struct {
	// ! required fields
	PlayerSeasonYear string `json:"seasonYear" jsonschema:"The season year of the player to compare (e.g. '2023-24')"`
	PlayerName       string `json:"playerName" jsonschema:"The name of the player to find similar players for (case-insensitive). The player must exist in the specified PlayerSeasonYear."`

	TopK               int    `json:"topK,omitempty" jsonschema:"The number of similar players to return (default: 5, maximum: 10)"`
	SeasonYearMoreThan string `json:"seasonYearMoreThan,omitempty" jsonschema:"Optional: Start season year for search (e.g., '2020-21'). Defaults to one year before PlayerSeasonYear. Earliest available is '2001-02'."`
	SeasonLessThan     string `json:"seasonYearLessThan,omitempty" jsonschema:"Optional: End season year for search (e.g., '2022-23'). Defaults to PlayerSeasonYear. Latest available is the current season."`
}

type DraftCombineSimilarityRecord struct {
	types.DraftCombineStatsRecord
	SimilarityScore float64 `json:"similarityScore" jsonschema:"This represents the similarity of the players' draft combine results, calculated using the root mean square error. The closer the value is to 0, the higher the similarity."`
}

type GetCombineSimilarityResult struct {
	SimilarPlayers []DraftCombineSimilarityRecord `json:"similarPlayers" jsonschema:"List of players with the most similar draft combine results to the specified player, along with their similarity score."`
}
