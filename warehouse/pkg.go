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
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	if mgr.dbs == nil {
		mgr.dbs = make(map[string]stores.Modules)
	}
	mgr.dbs[name] = conn
}

func GetRoleList(ctx context.Context, params map[string]interface{}, opts *Options) <-chan *stores.Data {
	out := make(chan *stores.Data, 1)

	go func() {
		defer close(out)

		// if we have a specific database to call
		if opts != nil && opts.PreferredDatabase != "" {
			db := mgr.dbs[opts.PreferredDatabase]
			data, err := db.GetRoleList(ctx, params)
			if err != nil {
				log.WithError(err).Error("failed to get roles for " + fmt.Sprintf("%T", db))
				return
			}
			out <- data
			return
		}

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		wg := sync.WaitGroup{}

		// fan out to all connected databases and wait for first return
		for name, conn := range mgr.dbs {
			wg.Add(1)
			log.Info("calling database: " + name)
			go func(db stores.Modules) {
				defer wg.Done()
				data, err := db.GetRoleList(ctx, params)
				if err != nil {
					log.WithError(err).Error("failed to get roles for " + fmt.Sprintf("%T", db))
					return
				}
				out <- data
				cancel()
			}(conn)
		}

		// all methods should properly exit when context is done
		wg.Wait()
	}()

	// return results
	return out
}
