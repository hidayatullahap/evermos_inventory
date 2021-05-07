package pkg

import (
	"math/rand"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"
)

var ulidSync = sync.Mutex{}

func NewULID() string {
	ulidSync.Lock()
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	ulidSync.Unlock()
	return id.String()
}
