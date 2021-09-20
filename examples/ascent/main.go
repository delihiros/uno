package main

import (
	"log"
	"uno/pkg/analysis/maps"
	"uno/pkg/jsonutil"

	"uno/pkg/client"
)

func main() {
	ascent := maps.NewAscent()
	c := client.New()
	match, err := c.GetMatchByID("3ea94c2f-3781-42cc-8862-b631c6756692")
	if err != nil {
		panic(err)
	}
	for r, round := range match.Rounds {
		for _, status := range round.PlayerStats {
			for _, event := range status.KillEvents {
				victimLocation := event.VictimDeathLocation
				killerLocation := event.FindKillerLocation()
				if killerLocation != nil {
					ascent.DrawCircle(float64(victimLocation.X), float64(victimLocation.Y), 3, 1, 0, 0)
					ascent.DrawCircle(float64(killerLocation.X), float64(killerLocation.Y), 3, 0, 0, 1)
					ascent.DrawLine(float64(victimLocation.X), float64(victimLocation.Y), float64(killerLocation.X), float64(killerLocation.Y), 2, 0, 0.5, 0.5)
				} else {
					s, err := jsonutil.FormatJSON(event, true)
					if err != nil {
						panic(err)
					}
					log.Println(r, s)
				}
			}
		}
	}
	ascent.SaveImage("death.png")
}
