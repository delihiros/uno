package entities

type MMRHistory struct {
	Currenttier         int    `json:"currenttier"`
	Currenttierpatched  string `json:"currenttierpatched"`
	RankingInTier       int    `json:"ranking_in_tier"`
	MmrChangeToLastGame int    `json:"mmr_change_to_last_game"`
	Elo                 int    `json:"elo"`
	Date                string `json:"date"`
	DateRaw             int64  `json:"date_raw"`
}
