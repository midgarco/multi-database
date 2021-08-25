package warehouse

import (
	"sync"

	"github.com/apex/log"
	"github.com/midgarco/multi-database/stores"
)

type manager struct {
	dbs map[string]stores.Interface
	mu  sync.Mutex
}

//
func (m *manager) GetConnections(connId int, opts *Options) []stores.Interface {
	// don't need the mutex here since we utilize
	// calls that already handle it

	list := []stores.Interface{}
	if opts != nil && len(opts.PreferredDatabase) > 0 {
		for _, pref := range opts.PreferredDatabase {
			dbint := m.GetDatabase(connId, pref.String())
			list = append(list, dbint)
		}
	} else {
		list = m.GetAllConnections()
	}
	return list
}

//
func (m *manager) GetAllConnections() []stores.Interface {
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
func (m *manager) GetDatabase(connId int, name string) stores.Module {
	m.mu.Lock()
	defer m.mu.Unlock()

	dbint := m.dbs[name]

	// if this is a cached type database, pull out the specific connection
	if dbint.GetModuleType() == stores.ModuleType_Cache {
		if _, ok := dbint.(stores.Cacher); !ok {
			return nil
		}
		db, err := dbint.(stores.Cacher).GetConnection("connection_one")
		if err != nil {
			log.WithError(err).Error("failed to get name connection")
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
