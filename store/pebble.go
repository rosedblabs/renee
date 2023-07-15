package store

import (
	"github.com/cockroachdb/pebble"
)

type PebbleStore struct {
	db      *pebble.DB
	options Options
}

func newPebbleStore(options Options) (Store, error) {
	pebbleOpts := new(pebble.Options)
	pebbleOpts.BytesPerSync = int(options.BytesPerSync)
	db, err := pebble.Open(options.DirPath, pebbleOpts)
	if err != nil {
		return nil, err
	}
	return &PebbleStore{db: db, options: options}, nil
}

func (ps *PebbleStore) Get(key []byte) ([]byte, error) {
	value, closer, err := ps.db.Get(key)
	if closer != nil {
		// handle the closer? todo
	}
	return value, err
}

func (ps *PebbleStore) Put(key []byte, value []byte) error {
	return ps.db.Set(key, value, &pebble.WriteOptions{Sync: ps.options.Sync})
}

func (ps *PebbleStore) Delete(key []byte) error {
	return ps.db.Delete(key, &pebble.WriteOptions{Sync: ps.options.Sync})
}

func (ps *PebbleStore) Close() error {
	return ps.db.Close()
}
