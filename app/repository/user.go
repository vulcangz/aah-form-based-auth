package repository

import (
	"aah-form-based-auth/app/models"
)

// UserEvent User Event struct
type UserEvent struct {
	User  *models.User `json:"user"`
	Roles []string     `json:"roles"`
	Perms []string     `json:"perms"`
}

// UpdateUserArgs User信息更新参数
type UpdateUserArgs struct {
	FirstName   string
	LastName    string
	Email       string
	Password    string
	IsLocked    bool
	Roles       []string
	Permissions []string
}

// UserRepository 用户存储库
type UserRepository interface {
	GetAllUsers() ([]*models.User, error)
	GetUserByID(string) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	CreateUser(*models.User) error
	Exist(email string) (bool, error)
	UpdateUser(email string, args UpdateUserArgs) error
}
