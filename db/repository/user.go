package db

import (
	"fmt"
)

type UserRepository interface{
	CreateUser() error
}

type UserRepoImpl struct{
//   db *sql.DB
}

func NewUserRepo() UserRepository {
	return &UserRepoImpl{}
}

func (r *UserRepoImpl) CreateUser() error {
	fmt.Println("creating user from repo")
	return nil
}