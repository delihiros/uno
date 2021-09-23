package main

import (
	"log"

	"github.com/delihiros/uno/pkg/database"
	"github.com/delihiros/uno/pkg/jsonutil"
)

func main() {
	db, err := database.Get()
	if err != nil {
		panic(err)
	}
	ms, err := db.ListMatch()
	if err != nil {
		panic(err)
	}
	for _, m := range ms {
		jsonutil.PrintJSON(m.Metadata, false)
	}
	log.Println(len(ms))
}
