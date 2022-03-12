package queue

import "aah-form-based-auth/app/models"

// UserEvent User Event struct
type UserEvent struct {
	User  *models.User `json:"user"`
	Roles []string     `json:"roles"`
	Perms []string     `json:"perms"`
}
