package entities

type LeaderboardPlayer struct {
	PlayerCardID    string `json:"PlayerCardID"`
	TitleID         string `json:"TitleID"`
	IsBanned        bool   `json:"IsBanned"`
	IsAnonymized    bool   `json:"IsAnonymized"`
	Puuid           string `json:"puuid"`
	GameName        string `json:"gameName"`
	TagLine         string `json:"tagLine"`
	LeaderboardRank int    `json:"leaderboardRank"`
	RankedRating    int    `json:"rankedRating"`
	NumberOfWins    int    `json:"numberOfWins"`
	CompetitiveTier int    `json:"competitiveTier"`
}
