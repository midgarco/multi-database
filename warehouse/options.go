package warehouse

import "github.com/midgarco/multi-database/stores"

type Options struct {
	PreferredDatabase []stores.DatabaseType
}

//
func (o *Options) AddPreferredDatabase(d ...stores.DatabaseType) {
	o.PreferredDatabase = append(o.PreferredDatabase, d...)
}
