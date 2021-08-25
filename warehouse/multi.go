package warehouse

import (
	"context"
	"fmt"

	"github.com/apex/log"
	"github.com/midgarco/multi-database/stores"
	"github.com/pkg/errors"
)

type CallerFunc func(ctx context.Context, params map[string]interface{}, dbint stores.Interface) *stores.Data

func Do(ctx context.Context, f CallerFunc, params map[string]interface{}, opts *Options) <-chan *stores.Data {
	out := make(chan *stores.Data, 1)

	// set options defaults
	if opts == nil {
		opts = &Options{}
	}

	go func() {
		defer close(out)

		// pull out the connection Id from the params
		connectionId, ok := params["connection_id"].(int)
		if !ok {
			out <- &stores.Data{Error: errors.New("missing connection_id parameter")}
			return
		}

		limit := 10000
		if _, ok := params["limit"]; ok {
			limit = params["limit"].(int)
		}

		// flag to know if we've already got data from a connection
		hadData := false

		// if we have preferred databases to call
	DB_CONNECTION_LOOP:
		for _, dbint := range mgr.GetConnections(connectionId, opts) {
			start := 0

			// last connection had data, so bail
			if hadData {
				break
			}

			// copy from the original map to the target map
			scopedParams := make(map[string]interface{}, len(params))
			for key, value := range params {
				scopedParams[key] = value
			}

			for {
				end := start + limit

				scopedParams["start"] = start
				scopedParams["end"] = end

				data := f(ctx, scopedParams, dbint)

				// flag the connection is good and has results
				if len(data.Results) > 0 {
					hadData = true
				}
				out <- data

				// we can clean these up and let the error bubble up
				if data.Error != nil {
					log.WithError(data.Error).Error("failed to get data for " + fmt.Sprintf("%T", dbint))
					continue DB_CONNECTION_LOOP
				}

				// check for data
				checkNext := false
				if len(data.Results) == limit && limit != 1 {
					checkNext = true
				}

				if !checkNext || (dbint.GetDatabaseType() == stores.DatabaseType_DynamoDB && data.NextToken == "") {
					log.WithFields(log.Fields{
						"connection_id": connectionId,
						"start":         start,
						"end":           end,
						"token":         scopedParams["next_token"],
						"database_type": fmt.Sprintf("%T", dbint),
					}).Debug("no more results")
					break
				}

				start = end
				scopedParams["next_token"] = data.NextToken
			}
		}
	}()

	return out
}
