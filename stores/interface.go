package stores

import (
	"context"
)

type (
	Interface interface {
		GetModuleType() ModuleType
		GetDatabaseType() DatabaseType
		Healthy() error
	}

	Cacher interface {
		GetConnection(key string) (Module, error)
	}

	Module interface {
		Interface

		GetEntityList(ctx context.Context, params map[string]interface{}) *Data
	}
)

type ModuleType string

var (
	ModuleType_DB    ModuleType = "db"
	ModuleType_Cache ModuleType = "cache"
)

//
func (mt ModuleType) String() string {
	return string(mt)
}

// ## DatabaseType
type DatabaseType string

var (
	DatabaseType_DynamoDB  DatabaseType = "dynamodb"
	DatabaseType_SqlServer DatabaseType = "mssql"
)

//
func (dt DatabaseType) String() string {
	return string(dt)
}
