package renee

import (
	"github.com/rosedblabs/renee/store"
	"time"
)

type Service struct {
	store   store.Store
	options Options
}

type DataType = uint8

const (
	String DataType = iota + 1
	Hash
	List
	Set
	ZSet
)

func New(options Options) (*Service, error) {
	backend, err := store.Open(store.Options{
		DirPath:      options.DirPath,
		Sync:         options.Sync,
		BytesPerSync: options.BytesPerSync,
	}, store.Type(options.BackendStorage))
	if err != nil {
		return nil, err
	}

	return &Service{store: backend, options: options}, nil
}

func (s *Service) Del(key []byte) error {
	return nil
}

func (s *Service) TTl(key []byte) (time.Duration, error) {
	return 0, nil
}
