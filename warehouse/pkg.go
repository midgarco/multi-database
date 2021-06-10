package warehouse

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/apex/log"
	"github.com/midgarco/multi-database/stores"
)

var mgr *manager = &manager{}

func AddConnection(name string, conn stores.Module) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	if mgr.dbs == nil {
		mgr.dbs = make(map[string]stores.Module)
	}
	mgr.dbs[name] = conn
}

func GetRoleList(ctx context.Context, params map[string]interface{}, opts *Options) (<-chan *stores.Data, error) {
	out := make(chan *stores.Data, 1)

	// pull out the client id from the params
	clientId, ok := params["client_id"].(int)
	if !ok {
		return nil, errors.New("missing client_id parameter")
	}
	limit := 10000

	start := 0

	go func() {
		defer close(out)

		// if we have a specific database to call
		if opts != nil && opts.PreferredDatabase != "" {
			dbint := mgr.GetPreferredDatabase(clientId, opts.PreferredDatabase)

			// adjust to the correct service interface
			db, ok := dbint.(stores.UserManager)
			if !ok {
				err := errors.New("preferred db does not satisfy the UserManager interface")
				out <- &stores.Data{Error: err}
				log.WithError(err).Error("failed to get preferred database")
				return
			}

			for {
				end := start + limit

				params["start"] = start
				params["end"] = end

				data, err := db.GetRoleList(ctx, params)
				if err != nil {
					log.WithError(err).Error("failed to get roles for " + fmt.Sprintf("%T", db))
					break
				}
				out <- data

				// check for data
				hasData := false
				if len(data.Results) > 0 {
					hasData = true
				}

				if !hasData {
					log.WithFields(log.Fields{
						"client_id": clientId,
						"start":     start,
						"end":       end,
					}).Debug("no results")
					break
				}

				if data.NextToken != "" {
					params["next_token"] = data.NextToken
				}

				start = end
			}

			return
		}

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		wg := sync.WaitGroup{}

		// fan out to all connected databases and wait for first return
		for _, conn := range mgr.GetAllConnections(clientId) {
			wg.Add(1)
			go func(dbint stores.Interface) {
				defer wg.Done()
				defer cancel()

				// adjust to the correct service interface
				db, ok := dbint.(stores.UserManager)
				if !ok {
					return
				}

				for {
					end := start + limit

					params["start"] = start
					params["end"] = end

					// get the data
					data, err := db.GetRoleList(ctx, params)
					if err != nil {
						log.WithError(err).Error("failed to get roles for " + fmt.Sprintf("%T", db))
						break
					}
					out <- data

					// check for data
					hasData := false
					if data != nil && len(data.Results) > 0 {
						hasData = true

						if data.NextToken != "" {
							params["next_token"] = data.NextToken
						}
					}

					if !hasData {
						log.WithFields(log.Fields{
							"client_id": clientId,
							"start":     start,
							"end":       end,
						}).Debug("no results")
						break
					}

					start = end
				}
			}(conn)
		}

		// all methods should properly exit when context is done
		wg.Wait()
	}()

	// return result channel
	return out, nil
}
