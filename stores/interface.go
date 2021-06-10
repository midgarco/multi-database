package stores

import (
	"context"
)

type (
	Interface interface {
		GetModuleType() ModuleType
		Healthy() error
	}

	Cacher interface {
		GetConnection(key string) (Module, error)
	}

	Module interface {
		Interface
		UserManager
	}

	UserManager interface {
		GetRoleList(ctx context.Context, params map[string]interface{}) (*Data, error)
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
