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
func (db *DynamoDB) GetRoleList(ctx context.Context, params map[string]interface{}) (*Data, error) {
	log.Info("dynamo get role list")

	var data Data
	var err error

	done := make(chan bool)
	go func() {
		// run the query information
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
		}
		done <- true
		close(done)
	}()

	for {
		select {
		case <-ctx.Done():
			log.Info("dynamo context done")
			return nil, nil
		case <-done:
			log.Info("dynamo role list results")
			return &data, err
		}
	}
}
