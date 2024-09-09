package repository

import (
	"database/sql"
	"log"

	"github.com/lauragrassig/golang-api/internal/schema"
)

type UserRepository struct{
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) GetUser(id int) (* schema.User, error) {
	r, err := ur.db.Query("SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		log.Printf("error selecting user on GetUser: %v\n", err)
		return nil, err
	}

	defer r.Close()

	var user schema.User

	for r.Next() {
		if err := r.Scan(&user.Id, &user.FirstName, &user.LastName, &user.BirthDate, &user.DriverLicense, &user.CreatedAt); err != nil {
			log.Printf("error scanning user on GetUser %v\n", err)
			return nil, err
		}
	}

	return &user, nil
}

func (ur *UserRepository) GetUsers() ([]*schema.User, error) {
	r, err := ur.db.Query("SELECT * FROM users ORDER BY id")

	if err != nil {
		log.Printf("error selecting users on GetUsers: %v\n", err)
		return nil, err
	}

	defer r.Close()

	var users []*schema.User

	for r.Next() {
		var user schema.User
		if err := r.Scan(&user.Id, &user.FirstName, &user.LastName, &user.BirthDate, &user.DriverLicense, &user.CreatedAt); err != nil {
			log.Printf("error scanning users on GetUsers: %v\n", err)
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (ur *UserRepository) CreateUser(user *schema.User) error {
	_, err := ur.db.Exec("INSERT INTO users (first_name, last_name, birth_date, driver_license, luck_number, created_at) values ($1, $2, $3, $4, $5, $6)", user.FirstName, user.LastName, user.BirthDate, user.DriverLicense, user.LuckNumber, user.CreatedAt)

	if err != nil {
		log.Printf("error inserting a new user: %v\n", err)
		return err
	}

	return nil
}