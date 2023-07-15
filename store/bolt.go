package store

import "go.etcd.io/bbolt"

var defaultCF = []byte("default")

type BoltStore struct {
	db      *bbolt.DB
	options Options
}

func newBoltStore(options Options) (Store, error) {
	boltOpts := bbolt.DefaultOptions
	boltOpts.NoSync = !options.Sync
	db, err := bbolt.Open(options.DirPath, 0600, boltOpts)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(defaultCF)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &BoltStore{db: db, options: options}, nil
}

func (bs *BoltStore) Get(key []byte) ([]byte, error) {
	var value []byte
	err := bs.db.View(func(tx *bbolt.Tx) error {
		value = tx.Bucket(defaultCF).Get(key)
		return nil
	})
	return value, err
}

func (bs *BoltStore) Put(key []byte, value []byte) error {
	return bs.db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket(defaultCF).Put(key, value)
	})
}

func (bs *BoltStore) Delete(key []byte) error {
	return bs.db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket(defaultCF).Delete(key)
	})
}

func (bs *BoltStore) Close() error {
	return bs.db.Close()
}
