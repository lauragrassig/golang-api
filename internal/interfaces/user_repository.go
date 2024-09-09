package interfaces

import "github.com/lauragrassig/golang-api/internal/schema"

type UserRepository interface {
	GetUser(int) (*schema.User, error)
	GetUsers() ([]*schema.User, error)
	CreateUser(*schema.User) error
}

