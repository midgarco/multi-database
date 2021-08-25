package mssql

import (
	"sync"

	"github.com/midgarco/multi-database/stores"
)

var _ stores.Cacher = (*Cache)(nil)

//
type Cache struct {
	sync.Map
}

//
func NewCache() *Cache {
	return new(Cache)
}

//
func (c *Cache) GetConnection(key string) (stores.Module, error) {
	if conn, ok := c.Load(key); ok {
		return conn.(stores.Module), nil
	}

	conn := &stores.SqlServer{}

	c.Store(key, conn)

	return conn, nil
}

//
func (c *Cache) Healthy() error {
	var err error
	c.Range(func(connId, conn interface{}) bool {
		db := conn.(*stores.SqlServer)
		err = db.Healthy()
		return err == nil
	})
	return err
}

//
func (*Cache) GetModuleType() stores.ModuleType {
	return stores.ModuleType_Cache
}
