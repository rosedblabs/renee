package store

type Store interface {
	Get([]byte) ([]byte, error)
	Put([]byte, []byte) error
	Delete([]byte) error
	Close() error
}

type Type uint8

const (
	Pebble Type = iota
	Badger
	Bolt
)

func Open(options Options, typ Type) (Store, error) {
	switch typ {
	case Pebble:
		return newPebbleStore(options)
	case Badger:
		return newBadgerStore(options)
	case Bolt:
		return newBoltStore(options)
	default:
		panic("unexpected store type")
	}
}
