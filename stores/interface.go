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
		// GetRoleById(ctx context.Context, roleId int) ([]interface{}, error)
		// GetRoleByName(ctx context.Context, roleName string) ([]interface{}, error)
		GetRoleList(ctx context.Context, params map[string]interface{}) ([]map[string]interface{}, error)
		// GetUserById(ctx context.Context, userId int) (map[string]interface{}, error)
		// GetUserList(ctx context.Context, params map[string]interface{}) ([]map[string]interface{}, error)
		// CreateUser(ctx context.Context, user map[string]interface{}) (map[string]interface{}, error)
		// UpdateUser(ctx context.Context, user map[string]interface{}) (map[string]interface{}, error)
	}
)
