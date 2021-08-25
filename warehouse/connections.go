package warehouse

import "github.com/midgarco/multi-database/stores"

// Manage active connections
func AddConnection(name string, conn stores.Interface) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	if mgr.dbs == nil {
		mgr.dbs = make(map[string]stores.Interface)
	}
	mgr.dbs[name] = conn
}

func GetConnection(connId int, name string) stores.Interface {
	return mgr.GetDatabase(connId, name)
}
