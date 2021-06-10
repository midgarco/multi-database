package stores

import "context"

type (
	Interface interface {
		Healthy() error
	}

	Modules interface {
		UserManager
	}

	UserManager interface {
		// GetRoleById(ctx context.Context, roleId int) (*Data, error)
		// GetRoleByName(ctx context.Context, roleName string) (*Data, error)
		GetRoleList(ctx context.Context, params map[string]interface{}) (*Data, error)
		// GetUserById(ctx context.Context, userId int) (*Data, error)
		// GetUserList(ctx context.Context, params map[string]interface{}) (*Data, error)
		// CreateUser(ctx context.Context, user map[string]interface{}) (*Data, error)
		// UpdateUser(ctx context.Context, user map[string]interface{}) (*Data, error)
	}
)
