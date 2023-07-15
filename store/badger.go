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

func (bs *BadgerStore) Get(key []byte) (value []byte, err error) {
	_ = bs.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil && err != badger.ErrKeyNotFound {
			return err
		}
		if err == badger.ErrKeyNotFound {
			return nil
		}
		value, err = item.ValueCopy(nil)
		return nil
	})
	return
}

func (bs *BadgerStore) Put(key []byte, value []byte) error {
	return bs.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
}

func (bs *BadgerStore) Delete(key []byte) error {
	return bs.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

func (bs *BadgerStore) Close() error {
	return bs.db.Close()
}
