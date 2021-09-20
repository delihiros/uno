package entities

type Player struct {
	Puuid              string `json:"puuid"`
	Name               string `json:"name"`
	Tag                string `json:"tag"`
	Team               string `json:"team"`
	Character          string `json:"character"`
	Currenttier        int    `json:"currenttier"`
	CurrenttierPatched string `json:"currenttier_patched"`
	PlayerCard         string `json:"player_card"`
	PlayerTitle        string `json:"player_title"`
	Stats              struct {
		Score   int `json:"score"`
		Kills   int `json:"kills"`
		Deaths  int `json:"deaths"`
		Assists int `json:"assists"`
	} `json:"stats"`
	AbilityCasts   AbilityCasts `json:"ability_casts"`
	DamageMade     int          `json:"damage_made"`
	DamageReceived int          `json:"damage_received"`
}
