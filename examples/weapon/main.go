package main

import (
	"log"
	"strings"

	"github.com/delihiros/uno/pkg/proxy"
)

const (
	matchID = "2aa59334-e53a-415b-bb3d-4832305ee7db"
)

func main() {
	p, err := proxy.New()
	if err != nil {
		panic(err)
	}
	m, err := p.GetMatchByID(matchID)
	if err != nil {
		panic(err)
	}
	for _, r := range m.Rounds {
		for _, s := range r.PlayerStats {
			for _, e := range s.KillEvents {
				reason := ""
				w, err := p.GetWeapon(e.DamageWeaponID)
				if err != nil {
					panic(err)
				}
				reason = w.DisplayName
				kill := strings.Join([]string{
					reason,
					e.KillerDisplayName,
					e.VictimDisplayName,
				}, " ")
				log.Println(kill)
			}
		}
	}
}
