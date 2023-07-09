package store

import "github.com/dgraph-io/badger/v4"

type BadgerStore struct {
	db      *badger.DB
	options Options
}

func newBadgerStore(options Options) (Store, error) {
	badgerOpts := badger.DefaultOptions(options.DirPath)
	badgerOpts.SyncWrites = options.Sync
	db, err := badger.Open(badgerOpts)
	if err != nil {
		return nil, err
	}
	return &BadgerStore{db: db, options: options}, nil
}

func (bs *BadgerStore) Get([]byte) ([]byte, error) {
	return nil, nil
}

func (bs *BadgerStore) Put([]byte, []byte) error {
	return nil
}

func (bs *BadgerStore) Delete([]byte) error {
	return nil
}

func (bs *BadgerStore) Close() error {
	return nil
}
