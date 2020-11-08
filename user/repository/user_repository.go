package repository

import (
	"github.com/ujinjinjin/services/user/factories"
)

type UserRepository struct {
	factory *factories.DbContextFactory
}

func NewUserRepository(factory *factories.DbContextFactory) *UserRepository {
	return &UserRepository{
		factory,
	}
}

func (s *UserRepository) Test() (string, error) {
	dbContext := s.factory.CreateUserDbContext()
	defer dbContext.Dispose()

	return dbContext.TestQuery()
}