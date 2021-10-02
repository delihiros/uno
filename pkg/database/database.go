package database

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/delihiros/shockv/pkg/client"

	"github.com/delihiros/uno/pkg/entities"
	"github.com/delihiros/uno/pkg/jsonutil"
)

var (
	matches  = "matches"
	contents = "contents"
	weapons  = "weapons"
	db       *Database
	gen      sync.Once
)

type Database struct {
	remote *client.Client
}

func Get(databaseURL string, port int) (*Database, error) {
	var err error
	err = nil
	gen.Do(func() {
		db = _new(databaseURL, port)
	})
	return db, err
}

func _new(databaseURL string, port int) *Database {
	return &Database{
		remote: client.New(databaseURL, port),
	}
}

func (db *Database) SetMatch(m *entities.Match) error {
	value, err := jsonutil.FormatJSON(m, false)
	if err != nil {
		return err
	}
	res, err := db.remote.Set(matches, m.Metadata.Matchid, value, 0)
	if err != nil {
		return err
	}
	if res.Status != http.StatusCreated {
		return fmt.Errorf(res.Message)
	}
	return nil
}

func (db *Database) ListMatch() ([]string, error) {
	res, err := db.remote.List(matches)
	if err != nil {
		return nil, err
	}
	if res.Status != http.StatusOK {
		return nil, fmt.Errorf(res.Message)
	}
	return res.Body, nil
}

func (db *Database) Match(matchID string) (*entities.Match, error) {
	res, err := db.remote.Get(matches, matchID)
	if err != nil {
		return nil, err
	}
	if res.Status != http.StatusOK {
		return nil, fmt.Errorf(res.Message)
	}
	var match entities.Match
	err = json.Unmarshal([]byte(res.Body), &match)
	if err != nil {
		return nil, err
	}
	return &match, err
}

func (db *Database) SetContent(c *entities.Content) error {
	s, err := jsonutil.FormatJSON(c, false)
	if err != nil {
		return err
	}
	res, err := db.remote.Set(contents, "content", s, 24*60*60)
	if err != nil {
		return err
	}
	if res.Status != http.StatusCreated {
		return fmt.Errorf(res.Message)
	}
	return nil
}

func (db *Database) Content() (*entities.Content, error) {
	res, err := db.remote.Get(contents, "content")
	if err != nil {
		return nil, err
	}
	if res.Status != http.StatusOK {
		return nil, fmt.Errorf(res.Message)
	}
	var content entities.Content
	err = json.Unmarshal([]byte(res.Body), &content)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (db *Database) SetWeapon(w *entities.Weapon) error {
	s, err := jsonutil.FormatJSON(w, false)
	if err != nil {
		return err
	}
	res, err := db.remote.Set(weapons, w.UUID, s, 0)
	if err != nil {
		return err
	}
	if res.Status != http.StatusCreated {
		return fmt.Errorf(res.Message)
	}
	return nil
}

func (db *Database) Weapon(uuid string) (*entities.Weapon, error) {
	if uuid == "Ultimate" || uuid == "Ability1" || uuid == "Ability2" {
		return &entities.Weapon{
			UUID:        uuid,
			DisplayName: uuid,
		}, nil
	}
	if uuid == "" {
		return &entities.Weapon{
			UUID:        "Spike",
			DisplayName: "Spike",
		}, nil
	}
	res, err := db.remote.Get(weapons, uuid)
	if err != nil {
		return nil, err
	}
	if res.Status != http.StatusOK {
		return nil, fmt.Errorf(res.Message)
	}
	var weapon entities.Weapon
	err = json.Unmarshal([]byte(res.Body), &weapon)
	if err != nil {
		return nil, err
	}
	return &weapon, nil
}
