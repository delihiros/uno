package maps

import "github.com/delihiros/uno/pkg/entities"

type HeatMap struct {
	*entities.GridMap
}

func NewHeatMap(gridSize, width, height int) *HeatMap {
	return &HeatMap{
		GridMap: entities.NewGridMap(gridSize, width, height),
	}
}
