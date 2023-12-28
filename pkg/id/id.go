package id

import (
	"time"

	"github.com/teris-io/shortid"
)

var generator *shortid.Shortid

func init() {
	generator = shortid.MustNew(0, shortid.DefaultABC, uint64(time.Now().Unix()))
}

// New returns a unique, non-sequential ID.
func New() string {
	return generator.MustGenerate()
}
