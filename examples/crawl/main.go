package main

import (
	"github.com/delihiros/uno/pkg/crawler"
	"github.com/delihiros/uno/pkg/proxy"
)

const (
	region          = "ap"
	seedAccountName = "bobobobobobobo"
	seedAccountTag  = "1212"
)

func main() {
	c, err := crawler.New()
	if err != nil {
		panic(err)
	}
	p, err := proxy.New()
	if err != nil {
		panic(err)
	}
	account, err := p.GetAccountByNameTag(seedAccountName, seedAccountTag)
	if err != nil {
		panic(err)
	}
	err = c.CrawlHistoryByPUUID(account.Region, account.Puuid, 100)
	if err != nil {
		panic(err)
	}
}
