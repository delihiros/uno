package entities

type KillEvent struct {
	KillTimeInRound       int              `json:"kill_time_in_round"`
	KillTimeInMatch       int              `json:"kill_time_in_match"`
	KillerPuuid           string           `json:"killer_puuid"`
	KillerDisplayName     string           `json:"killer_display_name"`
	KillerTeam            string           `json:"killer_team"`
	VictimPuuid           string           `json:"victim_puuid"`
	VictimDisplayName     string           `json:"victim_display_name"`
	VictimTeam            string           `json:"victim_team"`
	VictimDeathLocation   Location         `json:"victim_death_location"`
	DamageWeaponID        string           `json:"damage_weapon_id"`
	SecondaryFireMode     bool             `json:"secondary_fire_mode"`
	PlayerLocationsOnKill []PlayerLocation `json:"player_locations_on_kill"`
	Assistants            []Assistant      `json:"assistants"`
}

func (e *KillEvent) FindKillerLocation() *Location {
	killer := e.KillerPuuid
	for _, loc := range e.PlayerLocationsOnKill {
		if loc.PlayerPuuid == killer {
			return &loc.Location
		}
	}
	return nil
}

func (e *KillEvent) Equals(v *KillEvent) bool {
	return e.KillerPuuid == v.KillerPuuid && e.VictimPuuid == v.VictimPuuid && e.KillTimeInRound == v.KillTimeInRound
}
