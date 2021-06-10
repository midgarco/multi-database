package main

import (
	"context"

	"github.com/apex/log"
	"github.com/midgarco/multi-database/entity"
	"github.com/midgarco/multi-database/stores"
	"github.com/midgarco/multi-database/warehouse"
)

func main() {
	ddb := &stores.DynamoDB{}
	warehouse.AddConnection("dynamo", ddb)

	msql := &stores.SqlServer{}
	warehouse.AddConnection("mssql", msql)

	results, err := entity.GetList(context.Background(), 3)
	if err != nil {
		panic(err)
	}
	log.WithField("results", results).Info("query successful")
}
