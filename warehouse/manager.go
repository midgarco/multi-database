package warehouse

import (
	"sync"

	"github.com/apex/log"
	"github.com/midgarco/multi-database/stores"
)

type manager struct {
	dbs map[string]stores.Module
	mu  sync.Mutex
}

type Options struct {
	PreferredDatabase string
}

//
func (m *manager) GetAllConnections(clientId int) []stores.Interface {
	m.mu.Lock()
	defer m.mu.Unlock()

	list := []stores.Interface{}
	// loop over all connected dbs
	for _, conn := range m.dbs {
		// check for cached connections
		if _, ok := conn.(stores.Cacher); ok {
			cacheConn, _ := conn.(stores.Cacher).GetConnection("connection_one")
			conn, ok = cacheConn.(stores.Module)
			if !ok {
				log.Warn("cached connection did not implement stores.Module interface")
				continue
			}
		}
		list = append(list, conn)
	}

	return list
}

//
func (m *manager) GetPreferredDatabase(clientId int, preferred string) stores.Module {
	dbint := m.dbs[preferred]

	// if this is a cached type database, pull out the clientId specific connection
	if dbint.GetModuleType() == stores.ModuleType_Cache {
		if _, ok := dbint.(stores.Cacher); !ok {
			return nil
		}
		db, err := dbint.(stores.Cacher).GetConnection("connection_one")
		if err != nil {
			log.WithError(err).Error("failed to get preferred connection")
			return nil
		}
		return db.(stores.Module)
	}

	// global database connection
	db, ok := dbint.(stores.Module)
	if !ok {
		return nil
	}
	return db
}
