package entities

type MMRData struct {
	Name        string        `json:"name"`
	Tag         string        `json:"tag"`
	CurrentData CurrentSeason `json:"current_data"`
	BySeason    struct {
		E3A2 SeasonMMR `json:"e3a2"`
		E3A1 SeasonMMR `json:"e3a1"`
		E2A3 SeasonMMR `json:"e2a3"`
		E2A2 SeasonMMR `json:"e2a2"`
		E2A1 SeasonMMR `json:"e2a1"`
		E1A3 SeasonMMR `json:"e1a3"`
		E1A2 SeasonMMR `json:"e1a2"`
		E1A1 SeasonMMR `json:"e1a1"`
	} `json:"by_season"`
}
