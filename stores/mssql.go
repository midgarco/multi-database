package stores

import (
	"context"
	"errors"
	"time"

	"github.com/apex/log"
)

type SqlServer struct{}

func (db *SqlServer) GetRoleList(ctx context.Context, params map[string]interface{}) ([]map[string]interface{}, error) {
	log.Info("sql server get role list")

	var results []map[string]interface{}
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
			return results, err
		}
	}
}