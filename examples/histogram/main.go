package main

import (
	"github.com/delihiros/uno/pkg/database"
	"github.com/delihiros/uno/pkg/view"
)

func main() {
	db, err := database.Get()
	if err != nil {
		panic(err)
	}
	matches, err := db.ListMatch()
	averages := []float64{}
	for _, match := range matches {
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
