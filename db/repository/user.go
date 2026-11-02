package db

import (
	"fmt"
)

type UserRepository interface{
  	CreateUser() error
}

type UserRepo1 struct{
//   db *sql.DB
}

func NewUserRepo() UserRepository {
	return &UserRepo1{}
}

func (r *UserRepo1) CreateUser() error {
	fmt.Println("creating user from repo")
	return nil
}