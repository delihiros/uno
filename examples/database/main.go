package main

import (
	"io/ioutil"
	"log"

	"github.com/delihiros/uno/pkg/entities"

	"github.com/delihiros/uno/pkg/database"
	"github.com/delihiros/uno/pkg/jsonutil"
)

func main() {
	JustCount()
}

func JustCount() {
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

func CountMap() {
	db, err := database.Get()
	if err != nil {
		panic(err)
	}
	ms, err := db.ListMatch()
	if err != nil {
		panic(err)
	}
	maps := map[string][]*entities.Match{}
	for _, m := range ms {
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

func GenerateMatchJSON() {
	db, err := database.Get()
	if err != nil {
		panic(err)
	}
	ms, err := db.ListMatch()
	if err != nil {
		panic(err)
	}
	txt, err := jsonutil.FormatJSON(ms, true)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("matches.json", []byte(txt), 0666)
	if err != nil {
		panic(err)
	}
}
