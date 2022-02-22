package repository

import (
	"aah-form-based-auth/app/models"
)

// UserRepository 用户存储库
type UserRepository interface {
	GetUserByID(string) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	CreateUser(*models.User) error
}
