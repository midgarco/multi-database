package warehouse

import (
	"context"
	"fmt"
	"sync"

	"github.com/apex/log"
	"github.com/midgarco/multi-database/stores"
)

var mgr *manager = &manager{}

func AddConnection(name string, conn stores.Modules) {
	if mgr.dbs == nil {
		mgr.dbs = make(map[string]stores.Modules)
	}
	mgr.dbs[name] = conn
}

func GetRoleList(ctx context.Context, params map[string]interface{}, opts *Options) ([]map[string]interface{}, error) {
	// if we have a specific database to call
	if opts != nil && opts.PreferredDatabase != "" {
		db := mgr.dbs[opts.PreferredDatabase]
		return db.GetRoleList(ctx, params)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var results []map[string]interface{}
	var err error

	wg := sync.WaitGroup{}

	// fan out to all connected databases and wait for first return
	for name, conn := range mgr.dbs {
		wg.Add(1)
		log.Info("calling database: " + name)
		go func(db stores.Modules) {
			defer wg.Done()
			results, err = db.GetRoleList(ctx, params)
			if err != nil {
				log.WithError(err).Error("failed to get roles for " + fmt.Sprintf("%T", db))
				return
			}
			cancel()
		}(conn)
	}

	// all methods should properly exit when context is done
	wg.Wait()

	// return results
	return results, err
}
