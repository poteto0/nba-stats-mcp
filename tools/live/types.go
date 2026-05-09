package live

import "github.com/poteto0/go-nba-sdk/types"

type GetScoreBoardInput struct{}

type LiveScoreBoardRecord struct {
	Games []types.Game `json:"games" jsonschema:"The full scoreboard data structure as returned by the GNS API, containing all details about live games, teams, scores, time remaining, etc."`
}

type GetScoreBoardResult struct {
	Scoreboard LiveScoreBoardRecord `json:"scoreboard" jsonschema:"List of live games currently in progress, with detailed scoreboard information for each game"`
}

type GameSummaryRecord struct {
	GameId        string `json:"gameId" jsonschema:"Unique identifier for the game"`
	GameStatus    string `json:"gameStatus" jsonschema:"Current status of the game (e.g. 'In Progress', 'Final', 'Scheduled')"`
	HomeTeamName  string `json:"homeTeamName" jsonschema:"Name of the home team"`
	AwayTeamName  string `json:"awayTeamName" jsonschema:"Name of the away team"`
	HomeTeamScore int    `json:"homeTeamScore" jsonschema:"Current score of the home team"`
	AwayTeamScore int    `json:"awayTeamScore" jsonschema:"Current score of the away team"`
	GameClock     string `json:"gameClock" jsonschema:"Time remaining in the current period (e.g. '5:32')"`
	Period        int    `json:"period" jsonschema:"Current period of the game (e.g. 1 for first quarter, 2 for second quarter, etc.)"`
}

type LiveScoreBoardSummaryRecord struct {
	Games []GameSummaryRecord `json:"games" jsonschema:"List of live games currently in progress, with summary information for each game"`
}

type GetScoreBoardSummaryResult struct {
	Scoreboard LiveScoreBoardSummaryRecord `json:"scoreboard" jsonschema:"List of live games currently in progress, with summary information for each game"`
}

type GameLeaderRecord struct {
	GameId        string            `json:"gameId" jsonschema:"Unique identifier for the game"`
	GameStatus    string            `json:"gameStatus" jsonschema:"Current status of the game (e.g. 'In Progress', 'Final', 'Scheduled')"`
	HomeTeamName  string            `json:"homeTeamName" jsonschema:"Name of the home team"`
	AwayTeamName  string            `json:"awayTeamName" jsonschema:"Name of the away team"`
	HomeTeamScore int               `json:"homeTeamScore" jsonschema:"Current score of the home team"`
	AwayTeamScore int               `json:"awayTeamScore" jsonschema:"Current score of the away team"`
	GameClock     string            `json:"gameClock" jsonschema:"Time remaining in the current period (e.g. '5:32')"`
	Period        int               `json:"period" jsonschema:"Current period of the game (e.g. 1 for first quarter, 2 for second quarter, etc.)"`
	GameLeaders   types.GameLeaders `json:"gameLeaders" jsonschema:"Current statistical leaders for the game, including points, rebounds, assists, etc. for both teams"`
}

type GetGameLeadersResult struct {
	Scoreboard []GameLeaderRecord `json:"scoreboard" jsonschema:"List of live games currently in progress, with summary information and current statistical leaders for each game"`
}

type BoxScoreInput struct {
	GameId string `json:"gameId" jsonschema:"Unique identifier for the game to retrieve the box score for"`
}

type BoxScoreRecord struct {
	Games types.Game `json:"games" jsonschema:"Box score data for the specified game, including detailed statistics for players and teams"`
}

type GetBoxScoreResult struct {
	BoxScore BoxScoreRecord `json:"boxScore" jsonschema:"Box score data for the specified game, including detailed statistics for players and teams"`
}

type PlayByPlayInput struct {
	GameId      string `json:"gameId" jsonschema:"Unique identifier for the game to retrieve the play-by-play data for"`
	TeamId      int    `json:"teamId,omitempty" jsonschema:"Optional unique identifier for a specific team to filter the play-by-play data (if not provided, returns play-by-play for both teams)"`
	ActionType  string `json:"actionType,omitempty" jsonschema:"Optional filter for the type of play-by-play actions to include (e.g. 'shot', 'rebound', 'assist', etc.)"`
	ShotResult  string `json:"shotResult,omitempty" jsonschema:"Optional filter for the result of shot actions to include (e.g. 'made', 'missed', etc.)"`
	Period      int    `json:"period,omitempty" jsonschema:"Optional filter for the period of the game to include play-by-play data for (e.g. 1 for first quarter, 2 for second quarter, etc.)"`
	IsFieldGoal bool   `json:"isFieldGoal,omitempty" jsonschema:"Optional filter to include only field goal actions in the play-by-play data"`
}

type PlayByPlaySummaryRecord struct {
	ActionNumber int    `json:"actionNumber" jsonschema:"Sequential number of the play in the game"`
	ActionType   string `json:"actionType" jsonschema:"Type of the play-by-play action (e.g. 'turnover', '2pt', 'block', etc.)"`
	Period       int    `json:"period" jsonschema:"Period of the game in which the play occurred"`
	Description  string `json:"description" jsonschema:"Text description of the play, including details about the action, players involved, and outcome"`
}

type GetPlayByPlaySummaryResult struct {
	PlayByPlay []PlayByPlaySummaryRecord `json:"playByPlay" jsonschema:"List of play-by-play actions for the specified game, in sequential order"`
}

type GetSpecificPlayByPlayDetailsInput struct {
	GameId        string `json:"gameId" jsonschema:"Unique identifier for the game to retrieve the play-by-play data for"`
	ActionNumbers []int  `json:"actionNumbers" jsonschema:"List of specific action numbers to retrieve from the play-by-play data (if not provided, returns all actions)"`
}

type PlayByPlayDetailsResult struct {
	Actions []types.Action `json:"actions" jsonschema:"List of play-by-play actions for the specified game, including all details about each action as returned by the GNS API"`
}
