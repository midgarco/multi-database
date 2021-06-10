package warehouse

import (
	"sync"

	"github.com/midgarco/multi-database/stores"
)

type manager struct {
	dbs map[string]stores.Modules
	mu  sync.Mutex
}

type Options struct {
	PreferredDatabase string
}
