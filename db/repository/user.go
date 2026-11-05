package db

import (
	"auth-service/db/models"
	"database/sql"
	"fmt"
)

type UserRepository interface{
  	CreateUser(name, email, hashedPassword string) *models.User
	FindByEmail(email string) *models.User
}

type UserRepo1 struct{
    db *sql.DB
}

func NewUserRepo(_db *sql.DB) UserRepository {
	return &UserRepo1{db: _db}
}

func (r *UserRepo1) CreateUser(name, email, hashedPassword string) *models.User {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, name, email, hashedPassword)
	if err != nil {
		fmt.Println(err)
	}
	return &models.User{Name: name, Email: email}
}

func (r *UserRepo1) FindByEmail(email string) *models.User {
	row := r.db.QueryRow("SELECT name, email FROM users WHERE email = ?", email)
	u := &models.User{}
	row.Scan(&u.Name, &u.Email, &u.Password)
	return u
}