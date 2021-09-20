package entities

type PlayerStatus struct {
	AbilityCasts      AbilityCasts  `json:"ability_casts"`
	PlayerPuuid       string        `json:"player_puuid"`
	PlayerDisplayName string        `json:"player_display_name"`
	PlayerTeam        string        `json:"player_team"`
	DamageEvents      []DamageEvent `json:"damage_events"`
	Damage            int           `json:"damage"`
	Bodyshots         int           `json:"bodyshots"`
	Headshots         int           `json:"headshots"`
	Legshots          int           `json:"legshots"`
	KillEvents        []KillEvent   `json:"kill_events"`
	Kills             int           `json:"kills"`
}
