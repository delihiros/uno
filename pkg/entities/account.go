package entities

type Account struct {
	Puuid        string `json:"puuid"`
	Region       string `json:"region"`
	AccountLevel int    `json:"account_level"`
	Name         string `json:"name"`
	Tag          string `json:"tag"`
}
