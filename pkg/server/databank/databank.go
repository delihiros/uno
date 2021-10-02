package databank

import (
	"sync"

	"github.com/delihiros/uno/pkg/analysis/maps"

	"github.com/delihiros/uno/pkg/proxy"
)

var (
	db  *Databank
	gen sync.Once
)

type Databank struct {
	p        *proxy.Proxy
	heatmaps map[string]*maps.HeatMap
}

func Get(databaseURL string, port int) (*Databank, error) {
	var err error
	// TODO
	gen.Do(func() {
		db, err = _new(databaseURL, port)
		if err != nil {
			panic(err)
		}
	})
	return db, err
}

func _new(databaseURL string, port int) (*Databank, error) {
	p, err := proxy.New(databaseURL, port)
	if err != nil {
		return nil, err
	}
	db := &Databank{p: p}
	db.heatmaps, err = db.generateGrids()
	if err != nil {
		return nil, err
	}
	return db, nil
}
