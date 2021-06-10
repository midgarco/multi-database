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

	out := warehouse.GetRoleList(context.Background(), nil, nil)

	for data := range out {
		fmt.Println(data)
	}
	return results, errs.ErrorOrNil()
}
