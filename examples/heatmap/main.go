package main

import (
	"github.com/delihiros/uno/pkg/analysis/maps"
	"github.com/delihiros/uno/pkg/database"
	"github.com/delihiros/uno/pkg/entities"
	"github.com/delihiros/uno/pkg/view"
	"github.com/thoas/go-funk"
)

const (
	mapName = "Ascent"
)

func main() {
	m, err := entities.NewMap(mapName)
	if err != nil {
		panic(err)
	}
	mv, err := view.NewMapVisualizer(m)
	if err != nil {
		panic(err)
	}
	hm := maps.NewHeatMap(8, mv.Width, mv.Height)

	db, err := database.Get()
	if err != nil {
		panic(err)
	}
	matches, err := db.ListMatch()
	if err != nil {
		panic(err)
	}

	matches = funk.Filter(matches, func(match *entities.Match) bool {
		return match.Metadata.Map == mapName
	}).([]*entities.Match)

	for _, match := range matches {

		for _, r := range match.Rounds {
			for _, s := range r.PlayerStats {
				for _, k := range s.KillEvents {
					// TODO: Scale
					x, y := m.Scale(mv.Width, mv.Height, k.VictimDeathLocation.X, k.VictimDeathLocation.Y)
					gx, gy := hm.BelongsTo(x, y)
					g := hm.Get(gx, gy)
					g.Value++
				}
			}
		}
	}

	for _, i := range hm.Grids {
		for _, j := range i {
			if j.Value > 0 {
				mv.DrawRectangle(float64(j.X*hm.GridSize), float64(j.Y*hm.GridSize), float64(hm.GridSize), float64(hm.GridSize), 1, 0, 0, j.Value/4, true)
			}
		}
	}
	mv.SaveImage("heatmap.png")
}
