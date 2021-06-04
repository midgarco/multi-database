package warehouse

import "github.com/midgarco/multi-database/stores"

type manager struct {
	dbs map[string]stores.Modules
}

type Options struct {
	PreferredDatabase string
}
