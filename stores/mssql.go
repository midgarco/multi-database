package stores

import (
	"context"
	"errors"
	"time"

	"github.com/apex/log"
)

type SqlServer struct {
}

// Healthy check method to be compatible with the store.Interface
func (db *SqlServer) Healthy() error {
	// check the db connection
	return nil
}

//
func (db *SqlServer) GetModuleType() ModuleType {
	return ModuleType_DB
}

func (db *SqlServer) GetRoleList(ctx context.Context, params map[string]interface{}) (*Data, error) {
	log.Info("sql server get role list")

	var data Data
	var err error

	// testing errors
	err = errors.New("failure somewhere")

	done := make(chan bool)
	go func() {
		// run the query information
		for i := 0; i < 1; i++ {
			time.Sleep(time.Second)
		}
		done <- true
		close(done)
	}()

	for {
		select {
		case <-ctx.Done():
			log.Info("sql server context done")
			return nil, nil
		case <-done:
			log.Info("sql server role list results")
			return &data, err
		}
	}
}
