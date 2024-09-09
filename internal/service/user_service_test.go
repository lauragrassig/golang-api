package service

import (
	"errors"
	"testing"

	"github.com/lauragrassig/golang-api/internal/repository"
	"github.com/lauragrassig/golang-api/internal/schema"
)

func TestGenerateLuckNumber(t *testing.T) {
	tcases := map [string]struct{
		user schema.User
		expectedResult int
	}{
		"valid result": { 
				user: schema.User{
				FirstName: "Patrick",
				LastName: "Quengo",
				DriverLicense: 123456,
			},
			expectedResult: 22,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			generateUserLuckNumber(&tcase.user)
			if tcase.user.LuckNumber != tcase.expectedResult {
				t.Errorf("error validating luck number. expected %d, returned %d", tcase.expectedResult, tcase.user.LuckNumber)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	mockedUserRepository := repository.NewMockUserRepository()
	userService := NewUserService(mockedUserRepository)

	tcases := map [string]struct{
		user schema.CreateUserRequest
		expectedError error
	}{
		"valid user": { 
				user: schema.CreateUserRequest{
				FirstName: "Patrick",
				LastName: "Quengo",
				DriverLicense: 123456,
			},
			expectedError: nil,
		},
		"invalid first name ": { 
				user: schema.CreateUserRequest{
				FirstName: "",
				LastName: "Quengo",
				DriverLicense: 123456,
			},
			expectedError: ErrFirstNameEmpty,
	},
	"invalid last name ": { 
		user: schema.CreateUserRequest{
		FirstName: "Patrick",
		LastName: "",
		DriverLicense: 123456,
	},
	expectedError: ErrLastNameEmpty,
	},
	"invalid driver license": { 
		user: schema.CreateUserRequest{
		FirstName: "Patrick",
		LastName: "Quengo",
	},
	expectedError: ErrDriverLicenseEmpty,
	},
}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			err := userService.CreateUser(&tcase.user)

			if !errors.Is(err, tcase.expectedError) {
				t.Errorf( "error validating create user. Expected: %v, error received: %v", tcase.expectedError, err)
			}
		// assert.Equalf(t, err, tcase.expectedError, "error validating create user. Expected: %v, error received: %v", tcase.expectedError, err)
		})
	}
}