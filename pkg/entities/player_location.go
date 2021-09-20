package entities

type PlayerLocation struct {
	Location          Location `json:"location"`
	PlayerPuuid       string   `json:"player_puuid"`
	PlayerDisplayName string   `json:"player_display_name"`
	PlayerTeam        string   `json:"player_team"`
}
