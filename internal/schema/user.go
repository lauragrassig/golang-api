package schema

import "time"

type User struct {
	Id int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	DriverLicense int `json:"driver_license"`
	LuckNumber int `json:"luck_number"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserRequest struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	DriverLicense int `json:"driver_license"`
}