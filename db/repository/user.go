package db

import (
	"auth-service/db/models"
	"database/sql"
	"fmt"
)

type UserRepository interface{
  	CreateUser(email, hashedPassword string) *models.User
	FindByEmail(email string) *models.User
}

type UserRepo1 struct{
    db *sql.DB
}

func NewUserRepo(_db *sql.DB) UserRepository {
	return &UserRepo1{db: _db}
}

func (r *UserRepo1) CreateUser(email, hashedPassword string) *models.User {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	_, err := r.db.Exec(query, email, hashedPassword)
	if err != nil {
		fmt.Println(err)
	}
	return &models.User{Email: email}
}

func (r *UserRepo1) FindByEmail(email string) *models.User {
	row := r.db.QueryRow("SELECT id, email, password FROM users WHERE email = ?", email)
	u := &models.User{}
	if err := row.Scan(&u.Id, &u.Email, &u.Password); err != nil {
		return nil
	}
	return u
}