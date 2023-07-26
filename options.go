package renee

import "os"

type Options struct {
	// DirPath specifies the directory path where the WAL segment files will be stored.
	DirPath string

	// Sync is whether to synchronize writes through os buffer cache and down onto the actual disk.
	// Setting sync is required for durability of a single write operation, but also results in slower writes.
	//
	// If false, and the machine crashes, then some recent writes may be lost.
	// Note that if it is just the process that crashes (machine does not) then no writes will be lost.
	//
	// In other words, Sync being false has the same semantics as a write
	// system call. Sync being true means write followed by fsync.
	Sync bool

	// BytesPerSync specifies the number of bytes to write before calling fsync.
	BytesPerSync uint32

	// BackendStorage specifies the backend storage to use.
	BackendStorage BackendStorage
}

type BackendStorage uint8

const (
	Pebble BackendStorage = iota
	Badger
	Bolt
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

var DefaultOptions = Options{
	DirPath:        tempDBDir(),
	Sync:           false,
	BytesPerSync:   0,
	BackendStorage: Pebble,
}

func tempDBDir() string {
	dir, _ := os.MkdirTemp("", "renee-temp")
	return dir
}
