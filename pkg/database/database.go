package database

import (
	"bytes"
	"encoding/gob"
	"path/filepath"
	"time"
	"uno/pkg/entities"

	"github.com/dgraph-io/badger/v3"
)

var (
	databaseDirectory = "_databases"
	matchDirectory    = filepath.FromSlash(databaseDirectory + "/matches")
	contentDirectory  = filepath.FromSlash(databaseDirectory + "/content")
)

type Database struct {
	match   *badger.DB
	content *badger.DB
}

func New() (*Database, error) {
	match, err := badger.Open(badger.DefaultOptions(matchDirectory))
	content, err := badger.Open(badger.DefaultOptions(contentDirectory))
	if err != nil {
		return nil, err
	}
	return &Database{
		match:   match,
		content: content,
	}, nil
}

func (db *Database) SetMatch(m *entities.Match) error {
	id := []byte(m.Metadata.Matchid)
	bs, err := EncodeMatch(m)
	if err != nil {
		return err
	}
	return db.match.Update(func(txn *badger.Txn) error {
		return txn.Set(id, bs)
	})
}

func (db *Database) ListMatch() ([]*entities.Match, error) {
	var bss [][]byte
	err := db.match.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()
		var bs []byte
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			return item.Value(func(val []byte) error {
				bs = append([]byte{}, val...)
				bss = append(bss, bs)
				return nil
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	var ms []*entities.Match
	for _, bs := range bss {
		m, err := DecodeMatch(bs)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	return ms, nil
}

func (db *Database) Match(matchID string) (*entities.Match, error) {
	id := []byte(matchID)
	var bs []byte
	err := db.match.View(func(txn *badger.Txn) error {
		item, err := txn.Get(id)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			bs = append([]byte{}, val...)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return DecodeMatch(bs)
}

func EncodeMatch(m *entities.Match) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	err := gob.NewEncoder(buffer).Encode(m)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func DecodeMatch(bs []byte) (*entities.Match, error) {
	var m entities.Match
	buffer := bytes.NewBuffer(bs)
	err := gob.NewDecoder(buffer).Decode(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (db *Database) SetContent(c *entities.Content) error {
	id := []byte("content")
	bs, err := EncodeContent(c)
	if err != nil {
		return err
	}
	return db.content.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(id, bs).WithTTL(24 * time.Hour)
		return txn.SetEntry(e)
	})
}

func (db *Database) Content() (*entities.Content, error) {
	id := []byte("content")
	var bs []byte
	err := db.content.View(func(txn *badger.Txn) error {
		item, err := txn.Get(id)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			bs = append([]byte{}, val...)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return DecodeContent(bs)
}

func EncodeContent(c *entities.Content) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	err := gob.NewEncoder(buffer).Encode(c)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func DecodeContent(bs []byte) (*entities.Content, error) {
	var c entities.Content
	buffer := bytes.NewBuffer(bs)
	err := gob.NewDecoder(buffer).Decode(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
