package main

import (
	"github.com/delihiros/uno/pkg/crawler"
)

const (
	/*
		region          = "ap"
		seedAccountName = "bobobobobobobo"
		seedAccountTag  = "1212"
	*/
	region    = "eu"
	seedPUUID = "816b2f6e-daa6-5edb-bae1-21dd5ea8f6fa"
)

func main() {
	c, err := crawler.New()
	if err != nil {
		panic(err)
	}
	/*
		p, err := proxy.New()
		if err != nil {
			panic(err)
		}
		account, err := p.GetAccountByNameTag(seedAccountName, seedAccountTag)
		if err != nil {
			panic(err)
		}
	*/
	err = c.CrawlHistoryByPUUID(region, "816b2f6e-daa6-5edb-bae1-21dd5ea8f6fa", 100)
	if err != nil {
		panic(err)
	}
}
