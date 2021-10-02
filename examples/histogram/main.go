package main

import (
	"github.com/delihiros/uno/pkg/database"
	"github.com/delihiros/uno/pkg/view"
)

const (
	databaseURL = "http://localhost"
	port        = 8080
)

func main() {
	db, err := database.Get(databaseURL, port)
	if err != nil {
		panic(err)
	}
	matchIDs, err := db.ListMatch()
	averages := []float64{}
	for _, id := range matchIDs {
		match, err := db.Match(id)
		if err != nil {
			panic(err)
		}
		average, _, _ := match.Tier()
		averages = append(averages, average)
	}
	histogram, err := view.NewHistogram("Tier", averages, false)
	if err != nil {
		panic(err)
	}
	err = histogram.SaveImage("histogram.png")
	if err != nil {
		panic(err)
	}
}
