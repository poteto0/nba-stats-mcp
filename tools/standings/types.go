package standings

import "github.com/poteto0/go-nba-sdk/types"

type GetStandingsInput struct {
	Season     string `json:"season" jsonschema:"Season in the format YYYY-YY, default: 2025-26"`
	Limit      int    `json:"limit"  jsonschema:"Maximum number of records to return. Default is 30, must be 0 < limit <= 30. if conference filter is applied, limit is applied to each conference separately."`
	Conference string `json:"conference" jsonschema:"Optional filter for conference (East/West), default is no filter"`
}

type GetStandingsToolsResult struct {
	Standings []types.LeagueStandingsRecord `json:"standings" jsonschema:"List of league standings records, with monthly summarys and other stats"`
}

type GetStandingsSummaryRecord struct {
	TeamName   string `json:"teamName" jsonschema:"Team name"`
	Conference string `json:"conference" jsonschema:"Conference name (East/West)"`
	Wins       int    `json:"wins"`
	Losses     int    `json:"losses"`
}

type GetStandingsSummaryResult struct {
	Standings []GetStandingsSummaryRecord `json:"standings" jsonschema:"List of league standings records only having team name, conference, wins, and losses"`
}

type ClutchStandingsRecord struct {
	TeamName       string `json:"teamName" jsonschema:"Team name"`
	Conference     string `json:"conference" jsonschema:"Conference name (East/West)"`
	ThreePTSOrLess string `json:"threePTSOrLess" jsonschema:"Win-Loss record in games decided by 3 points or less (e.g. '15-8')"`
	OT             string `json:"ot"             jsonschema:"Win-Loss record in overtime games. critical for understanding performance in close games."`
	AheadAtHalf    string `json:"aheadAtHalf"   jsonschema:"Win-Loss record when leading at halftime"`
	BehindAtHalf   string `json:"behindAtHalf"  jsonschema:"Win-Loss record when trailing at halftime"`
	AheadAtThird   string `json:"aheadAtThird"  jsonschema:"Win-Loss record when leading after 3rd quarter"`
	BehindAtThird  string `json:"behindAtThird" jsonschema:"Win-Loss record when trailing after 3rd quarter"`
}

type ClutchStandingsResult struct {
	Standings []ClutchStandingsRecord `json:"standings" jsonschema:"List of league standings records with clutch performance stats"`
}

type MonthlyRecordEntry struct {
	Month  string `json:"month"  jsonschema:"Month abbreviation (e.g. 'oct', 'nov'). null if no games played that month"`
	Record string `json:"record" jsonschema:"Win-Loss record for the month (e.g. '12-3'). null if no games played"`
}

type MonthlyStandingsRecord struct {
	TeamName   string               `json:"teamName"   jsonschema:"Team name"`
	Conference string               `json:"conference" jsonschema:"Conference name (East/West)"`
	Monthly    []MonthlyRecordEntry `json:"monthly"    jsonschema:"Win-Loss record for each month of the season in chronological order. Months with no games are omitted"`
}

type MonthlyStandingsResult struct {
	Standings []MonthlyStandingsRecord `json:"standings" jsonschema:"List of league standings records with monthly performance stats"`
}
