package main

import (
	"fmt"
	"log"
	"uno/pkg/client"

	"github.com/thoas/go-funk"
)

const (
	matchID           = "2aa59334-e53a-415b-bb3d-4832305ee7db"
	playerDisplayName = "bobobobobobobo#1212"
)

func main() {
	c := client.New()
	match, err := c.GetMatchByID(matchID)
	if err != nil {
		log.Println(err)
	}
	for i, round := range match.Rounds {
		myStatus := funk.Filter(round.PlayerStats, func(stat client.PlayerStatus) bool {
			return stat.PlayerDisplayName == playerDisplayName
		}).([]client.PlayerStatus)[0]
		victims := funk.Reduce(myStatus.KillEvents, func(acc []string, event client.KillEvent) []string {
			acc = append(acc, event.VictimDisplayName)
			return acc
		}, []string{}).([]string)
		if len(victims) > 0 {
			fmt.Printf("In round %v: you killed %v\n", i+1, victims)
		}
	}
}
