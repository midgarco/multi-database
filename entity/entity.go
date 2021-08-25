package entity

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/midgarco/multi-database/stores"
	"github.com/midgarco/multi-database/warehouse"
)

func GetList(ctx context.Context, connectionId int) ([]map[string]interface{}, error) {
	results := []map[string]interface{}{}
	errs := &multierror.Error{}

	// build up parameter list
	params := make(map[string]interface{})
	params["connection_id"] = connectionId

	// set the warehouse preferred options
	opts := &warehouse.Options{}
	opts.AddPreferredDatabase(stores.DatabaseType_DynamoDB)

	out, err := warehouse.GetEntityList(context.Background(), params, opts)
	if err != nil {
		return nil, err
	}

	for data := range out {
		fmt.Println(data)
	}

	return results, errs.ErrorOrNil()
}
