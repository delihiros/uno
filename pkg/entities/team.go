package entities

type Team struct {
	HasWon     bool `json:"has_won"`
	RoundsWon  int  `json:"rounds_won"`
	RoundsLost int  `json:"rounds_lost"`
}
