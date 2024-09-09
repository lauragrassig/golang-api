package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/lauragrassig/golang-api/internal/interfaces"
	"github.com/lauragrassig/golang-api/internal/schema"
)

var (
	ErrFirstNameEmpty error = errors.New("first name can't be empty")
	ErrLastNameEmpty error = errors.New("last name can't be empty")
	ErrDriverLicenseEmpty error = errors.New("driver license can't be empty")
)

type UserService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetUser(id int) (*schema.User, error ) {
	return u.repo.GetUser(id)
}

func (u *UserService) GetUsers() ([]*schema.User, error) {
	return u.repo.GetUsers()
}

func (u *UserService) CreateUser(cur *schema.CreateUserRequest) error {
	if cur.FirstName == ""  {
		return ErrFirstNameEmpty
	}

	if cur.LastName == ""  {
		return ErrLastNameEmpty
	}

	if cur.DriverLicense == 0  {
		return ErrDriverLicenseEmpty
	}

	 user := schema.User{
		FirstName: cur.FirstName,
		LastName: cur.LastName,
		BirthDate: cur.BirthDate,
		DriverLicense: cur.DriverLicense,
		CreatedAt: time.Now(),
	 }

	 generateUserLuckNumber(&user)
	 return u.repo.CreateUser(&user)
}

func generateUserLuckNumber(user *schema.User) {
	luckNumber := len(user.FirstName + user.LastName + strconv.Itoa(user.DriverLicense)) + 3
	user.LuckNumber = luckNumber
}