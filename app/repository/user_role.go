package repository

import "aah-form-based-auth/app/models"

type UserRoleRepository interface {
	CreateUserRoles(roles ...*models.Role) error
	GetAllUserRoles() ([]*models.Role, error)
	CreateUserPermissions(perms ...*models.Permission) error
	GetAllUserPermissions() ([]*models.Permission, error)
}
