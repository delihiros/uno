package server

import (
	"net/http"
	"strconv"

	"github.com/delihiros/uno/pkg/server/databank"

	"github.com/delihiros/uno/pkg/server/protocols"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	db *databank.Databank
}

func New(databaseURL string, port int) (*Server, error) {
	db, err := databank.Get(databaseURL, port)
	if err != nil {
		return nil, err
	}
	return &Server{
		db: db,
	}, nil
}

func (s *Server) Execute(port int) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/heatmap", s.Heatmap)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))
}

// GET /heatmap?map={mapname}&x={int}&y={int}
func (s *Server) Heatmap(c echo.Context) error {
	var req protocols.HeatmapRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	grid, err := s.db.Grid(req.Map, req.X, req.Y)
	if err != nil {
		return err
	}

	pointed := &protocols.Grid{
		Location: &protocols.Point{
			X: grid.X * 4,
			Y: grid.Y * 4,
		},
		Value: grid.Value,
	}

	edges := []*protocols.Edge{}

	for _, e := range grid.To {
		edge := &protocols.Edge{
			To: &protocols.Point{
				X: e.To.X * 4,
				Y: e.To.Y * 4,
			},
			Weight: e.Weight,
		}
		edges = append(edges, edge)
	}

	res := &protocols.HeatmapResponse{
		GridSize: 4,
		Grid:     pointed,
		Edges:    edges,
	}
	return c.JSON(http.StatusOK, res)
}
