package main

import (
	"fmt"
	"log"

	"github.com/delihiros/uno/pkg/proxy"

	"github.com/delihiros/uno/pkg/client"
	"github.com/delihiros/uno/pkg/entities"
	"github.com/thoas/go-funk"
)

const (
	matchID           = "2aa59334-e53a-415b-bb3d-4832305ee7db"
	playerDisplayName = "bobobobobobobo#1212"
)

func main() {
	p, err := proxy.New()
	if err != nil {
		panic(err)
	}
	m, err := p.GetMatchByID(matchID)
	events, err := m.Rounds[0].KillEvents()
	if err != nil {
		panic(err)
	}
	for _, event := range events {
		log.Println(event.KillTimeInRound, event.VictimDisplayName, event.KillerDisplayName)
	}
}

func ListKills() {
	c := client.NewHenrikdevClient()
	match, err := c.GetMatchByID(matchID)
	if err != nil {
		log.Println(err)
	}
	for i, round := range match.Rounds {
		myStatus := funk.Filter(round.PlayerStats, func(stat entities.PlayerStatus) bool {
			return stat.PlayerDisplayName == playerDisplayName
		}).([]entities.PlayerStatus)[0]
		victims := funk.Reduce(myStatus.KillEvents, func(acc []string, event entities.KillEvent) []string {
			acc = append(acc, event.VictimDisplayName)
			return acc
		}, []string{}).([]string)
		if len(victims) > 0 {
			fmt.Printf("In round %v: you killed %v\n", i+1, victims)
		}
	}

}
