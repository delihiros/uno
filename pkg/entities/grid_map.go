package entities

type GridMap struct {
	GridSize int
	Grids    [][]*Grid
}

func NewGridMap(gridSize, width, height int) *GridMap {
	gm := &GridMap{
		GridSize: gridSize,
		Grids:    [][]*Grid{},
	}
	for x := 0; x < width/gridSize; x++ {
		row := []*Grid{}
		for y := 0; y < height/gridSize; y++ {
			row = append(row, NewGrid(x, y))
		}
		gm.Grids = append(gm.Grids, row)
	}
	return gm
}

func (g *GridMap) Get(x, y int) *Grid {
	return g.Grids[x][y]
}

func (g *GridMap) BelongsTo(x, y float64) (int, int) {
	return int(x) / g.GridSize, int(y) / g.GridSize
}
