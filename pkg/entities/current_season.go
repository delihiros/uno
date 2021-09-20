package entities

import "encoding/json"

type CurrentSeason struct {
	Currenttier          int    `json:"currenttier"`
	Currenttierpatched   string `json:"currenttierpatched"`
	RankingInTier        int    `json:"ranking_in_tier"`
	MmrChangeToLastGame  int    `json:"mmr_change_to_last_game"`
	Elo                  int    `json:"elo"`
	GamesNeededForRating int    `json:"games_needed_for_rating"`
}

// TODO: incorrect implementation
func (s *CurrentSeason) UnmarshalJSON(data []byte) error {
	var m map[string]*json.RawMessage
	return json.Unmarshal(data, &m)
}
