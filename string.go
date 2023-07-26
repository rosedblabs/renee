package renee

import (
	"encoding/binary"
	"time"
)

func (s *Service) Set(key []byte, value []byte, ttl time.Duration) error {
	// encode the value
	encodeValue := make([]byte, 1+8+len(value))
	encodeValue[0] = String
	var expired int64 = 0
	if ttl != time.Duration(0) {
		expired = time.Now().Add(ttl).UnixNano()
	}
	var index = 1
	binary.LittleEndian.PutUint64(encodeValue[index:index+8], uint64(expired))
	index += 8

	copy(encodeValue[index:], value)

	// put to backend storage engine
	return s.store.Put(key, encodeValue)
}

func (s *Service) Get(key []byte) ([]byte, error) {
	value, err := s.store.Get(key)
	if err != nil {
		return nil, err
	}
	if len(value) == 0 {
		return nil, ErrKeyNotFound
	}

	dataType := value[0]
	if dataType != String {

	}
	var index = 1
	expired := binary.LittleEndian.Uint64(value[index : index+8])
	index += 8
	if expired != 0 && int64(expired) <= time.Now().UnixNano() {
		return nil, ErrKeyNotFound
	}

	result := make([]byte, len(value)-9)
	copy(result, value[index:])
	return result, nil
}
