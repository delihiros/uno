package crawler

import (
	"log"
	"time"

	"github.com/delihiros/uno/pkg/proxy"
)

type Crawler struct {
	p       *proxy.Proxy
	visited map[string]bool
}

func New() (*Crawler, error) {
	p, err := proxy.New()
	if err != nil {
		return nil, err
	}
	return &Crawler{
		p:       p,
		visited: map[string]bool{},
	}, nil
}

func (c *Crawler) CrawlHistoryByPUUID(region, puuid string, depth int) error {
	log.Println(puuid, depth)
	if depth == 0 || c.alreadyVisited(puuid) {
		return nil
	}
	matches, err := c.p.GetMatchHistoryByPUUID(region, puuid, "competitive")
	if err != nil {
		return err
	}
	c.visit(puuid)
	for _, m := range matches {
		for _, p := range m.Players.AllPlayers {
			err = c.CrawlHistoryByPUUID(region, p.Puuid, depth-1)
			if err != nil {
				log.Println(err)
				// return err
			}
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}

func (c *Crawler) alreadyVisited(puuid string) bool {
	_, ok := c.visited[puuid]
	return ok
}

func (c *Crawler) visit(puuid string) {
	c.visited[puuid] = true
}
