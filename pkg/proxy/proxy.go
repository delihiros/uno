package proxy

import (
	"github.com/delihiros/uno/pkg/client"
	"github.com/delihiros/uno/pkg/database"
	"github.com/delihiros/uno/pkg/entities"
)

type Proxy struct {
	*client.Client
	db *database.Database
}

func New() (*Proxy, error) {
	db, err := database.New()
	if err != nil {
		return nil, err
	}
	return &Proxy{
		Client: client.New(),
		db:     db,
	}, nil
}

func (p *Proxy) GetMatchByID(matchID string) (*entities.Match, error) {
	m, err := p.db.Match(matchID)
	if err != nil {
		m, err = p.Client.GetMatchByID(matchID)
		err = p.db.SetMatch(m)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

func (p *Proxy) GetContent() (*entities.Content, error) {
	c, err := p.db.Content()
	if err != nil {
		c, err := p.Client.GetContent()
		err = p.db.SetContent(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (p *Proxy) GetMatchHistory(region string, name string, tag string, filter string) ([]*entities.Match, error) {
	matches, err := p.Client.GetMatchHistory(region, name, tag, filter)
	if err != nil {
		return nil, err
	}
	for _, m := range matches {
		err = p.db.SetMatch(m)
		if err != nil {
			return nil, err
		}
	}
	return matches, nil
}

func (p *Proxy) GetWeapon(uuid string) (*entities.Weapon, error) {
	weapon, err := p.db.Weapon(uuid)
	if err != nil {
		ws, err := p.Client.GetWeapons()
		if err != nil {
			return nil, err
		}
		for _, w := range ws {
			err = p.db.SetWeapon(w)
			if err != nil {
				return nil, err
			}
		}
		return p.db.Weapon(uuid)
	} else {
		return weapon, nil
	}
}
