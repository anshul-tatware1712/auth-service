package db

//facilitates dependency injection for repos
type Storage struct {
	UserRepository UserRepository
}

func NewStorage() *Storage {
	return &Storage{UserRepository: &UserRepoImpl{}}
}