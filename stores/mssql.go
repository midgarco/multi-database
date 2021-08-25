package stores

import (
	"context"
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

//
func (db *SqlServer) GetDatabaseType() DatabaseType {
	return DatabaseType_SqlServer
}

func (db *SqlServer) GetEntityList(ctx context.Context, params map[string]interface{}) *Data {
	log.Info("sql server get entity list")

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
			log.Info("sql server context done")
			break LOOP
		case <-done:
			log.Info("sql server role list results")
			break LOOP
		}
	}

	return &data
}
