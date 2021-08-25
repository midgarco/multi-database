package warehouse

import (
	"context"
	"errors"

	"github.com/apex/log"
	"github.com/midgarco/multi-database/stores"
)

var mgr *manager = &manager{}

func GetEntityList(ctx context.Context, params map[string]interface{}, opts *Options) (<-chan *stores.Data, error) {
	log.WithFields(log.Fields{
		"params":  params,
		"options": opts,
	}).Debug("calling get entity list")

	f := func(ctx context.Context, params map[string]interface{}, dbint stores.Interface) *stores.Data {
		// adjust to the correct service interface
		db, ok := dbint.(stores.Module)
		if !ok {
			err := errors.New("db does not satisfy the Module interface")
			return &stores.Data{Error: err}
		}

		return db.GetEntityList(ctx, params)
	}

	out := Do(ctx, f, params, opts)

	// return result channel
	return out, nil
}
