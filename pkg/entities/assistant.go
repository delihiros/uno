package entities

type Assistant struct {
	AssistantPuuid       string `json:"assistant_puuid"`
	AssistantDisplayName string `json:"assistant_display_name"`
	AssistantTeam        string `json:"assistant_team"`
}
