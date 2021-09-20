package entities

type SeasonMMR struct {
	Wins             int    `json:"wins"`
	NumberOfGames    int    `json:"number_of_games"`
	FinalRank        int    `json:"final_rank"`
	FinalRankPatched string `json:"final_rank_patched"`
	ActRankWins      []struct {
		PatchedTier string `json:"patched_tier"`
		Tier        int    `json:"tier"`
	} `json:"act_rank_wins"`
}
