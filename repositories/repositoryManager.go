package repositories

import "database/sql"

type RepositoryManager struct {
	CategoryRepository
	UsersRepository
}

// constructor
func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*NewCategoryRepository(dbHandler),
		*NewUsersRepository(dbHandler),
	}
}
