package entities

type DamageEvent struct {
	ReceiverPuuid       string `json:"receiver_puuid"`
	ReceiverDisplayName string `json:"receiver_display_name"`
	ReceiverTeam        string `json:"receiver_team"`
	Bodyshots           int    `json:"bodyshots"`
	Damage              int    `json:"damage"`
	Headshots           int    `json:"headshots"`
	Legshots            int    `json:"legshots"`
}
