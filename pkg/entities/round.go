package entities

type Round struct {
	WinningTeam string `json:"winning_team"`
	EndType     string `json:"end_type"`
	BombPlanted bool   `json:"bomb_planted"`
	BombDefused bool   `json:"bomb_defused"`
	PlantEvents struct {
		PlantLocation Location `json:"plant_location"`
		PlantedBy     struct {
			DisplayName string `json:"display_name"`
			Team        string `json:"team"`
		} `json:"planted_by"`
		PlantSide              string           `json:"plant_side"`
		PlantTimeInRound       int              `json:"plant_time_in_round"`
		PlayerLocationsOnPlant []PlayerLocation `json:"player_locations_on_plant"`
	} `json:"plant_events"`
	DefuseEvents struct {
		DefusedBy               SimplePlayer     `json:"defused_by"`
		DefuseLocation          Location         `json:"defuse_location"`
		DefuseTimeInRound       int              `json:"defuse_time_in_round"`
		PlayerLocationsOnDefuse []PlayerLocation `json:"player_locations_on_defuse"`
	} `json:"defuse_events"`
	PlayerStats []PlayerStatus `json:"player_stats"`
}
