package store

import (
	"github.com/cockroachdb/pebble"
)

type PebbleStore struct {
	db *pebble.DB
}

func newPebbleStore() (*Store, error) {
	return nil, nil
}

func (ps *PebbleStore) Get([]byte) ([]byte, error) {
	return nil, nil
}

func (ps *PebbleStore) Put([]byte, []byte) error {
	return nil
}

func (ps *PebbleStore) Delete([]byte) error {
	return nil
}

func (ps *PebbleStore) Commit() error {
	return nil
}

func (ps *PebbleStore) Close() error {
	return nil
}
