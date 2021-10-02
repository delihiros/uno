package main

import (
	"log"

	"github.com/delihiros/shockv/pkg/client"

	"github.com/delihiros/uno/pkg/entities"

	"github.com/delihiros/uno/pkg/database"
)

const (
	databaseURL = "http://localhost"
	port        = 8080
)

func main() {
	CountMap()
}

func RemoteList() {
	c := client.New("http://localhost", 8080)
	ids, err := c.List("matches")
	if err != nil {
		panic(err)
	}
	log.Println(ids)
}

func JustCount() {
	db, err := database.Get(databaseURL, port)
	if err != nil {
		panic(err)
	}
	ms, err := db.ListMatch()
	if err != nil {
		panic(err)
	}
	log.Println(len(ms))
}

func CountMap() {
	db, err := database.Get(databaseURL, port)
	if err != nil {
		panic(err)
	}
	ms, err := db.ListMatch()
	if err != nil {
		panic(err)
	}
	maps := map[string][]*entities.Match{}
	for _, mID := range ms {
		m, err := db.Match(mID)
		if err != nil {
			panic(err)
		}
		_, ok := maps[m.Metadata.Map]
		if !ok {
			maps[m.Metadata.Map] = []*entities.Match{}
		}
		maps[m.Metadata.Map] = append(maps[m.Metadata.Map], m)
	}
	for k, v := range maps {
		log.Println(k, len(v))
	}
}
