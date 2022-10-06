package core

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func MakeULID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
