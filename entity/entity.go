package entity

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/midgarco/multi-database/warehouse"
)

func GetList(ctx context.Context, clientId int) ([]map[string]interface{}, error) {
	results := []map[string]interface{}{}
	errs := &multierror.Error{}

	// build up parameter list
	params := make(map[string]interface{})
	params["client_id"] = clientId

	out, err := warehouse.GetRoleList(context.Background(), params, nil)
	if err != nil {
		return nil, err
	}

	for data := range out {
		fmt.Println(data)
	}
	return results, errs.ErrorOrNil()
}
