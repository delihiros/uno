package main

import (
	"uno/pkg/jsonutil"
	"uno/pkg/proxy"
)

var (
	matchID = "2aa59334-e53a-415b-bb3d-4832305ee7db"
)

func main() {
	p, err := proxy.New()
	if err != nil {
		panic(err)
	}
	for i := 0; i < 100; i++ {
		m, err := p.GetMatchByID(matchID)
		if err != nil {
			panic(err)
		}
		jsonutil.PrintJSON(m.Metadata, false)
	}
}
