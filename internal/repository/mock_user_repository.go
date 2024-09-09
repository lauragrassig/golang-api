package repository

import (
	"errors"

	"github.com/lauragrassig/golang-api/internal/schema"
)

type MockUserRepository struct{}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

func (ur *MockUserRepository) GetUser(id int) (*schema.User, error) {
	if id == 0 {
		return nil, errors.New("user not found")
	}

	return &schema.User{
		FirstName: "Mocked firstname",
		LastName: "Mocked lastname",
		DriverLicense: 123,
		LuckNumber: 777,
	}, nil
}

func (ur *MockUserRepository) GetUsers() ([]*schema.User, error) {
	return nil, nil
}

func (ur *MockUserRepository) CreateUser(user *schema.User) error {
	return nil
}