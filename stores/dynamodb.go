package stores

import (
	"context"
	"time"

	"github.com/apex/log"
)

type DynamoDB struct{}

//
func (db *DynamoDB) Healthy() error {
	// check the db connection
	return nil
}

//
func (db *DynamoDB) GetModuleType() ModuleType {
	return ModuleType_DB
}

//
func (db *DynamoDB) GetDatabaseType() DatabaseType {
	return DatabaseType_DynamoDB
}

//
func (db *DynamoDB) GetEntityList(ctx context.Context, params map[string]interface{}) *Data {
	log.Info("dynamo get entity list")

	var data Data

	done := make(chan bool)
	go func() {
		// run the query information
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
		}
		done <- true
		close(done)
	}()

LOOP:
	for {
		select {
		case <-ctx.Done():
			log.Info("dynamo context done")
			break LOOP
		case <-done:
			log.Info("dynamo role list results")
			break LOOP
		}
	}

	return &data
}
