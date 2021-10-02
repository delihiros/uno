package protocols

type HeatmapRequest struct {
	Map string `query:"map" json:"map"`
	X   int    `query:"x" json:"x"`
	Y   int    `query:"y" json:"y"`
}

type HeatmapResponse struct {
	MapURL   string  `json:"map_url"`
	GridSize int     `json:"grid_size"`
	Grid     *Grid   `json:"grid_pointed"`
	Edges    []*Edge `json:"grids"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Grid struct {
	Location *Point  `json:"location"`
	Value    float64 `json:"value"`
}

type Edge struct {
	To     *Point  `json:"to"`
	Weight float64 `json:"weight"`
}
