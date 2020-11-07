package repository

import (
	"github.com/ujinjinjin/user_service/factories"
)

type UserRepository struct {
	*factories.DbContextFactory
}

func NewUserRepository(factory *factories.DbContextFactory) *UserRepository {
	return &UserRepository{
		factory,
	}
}

func (s *UserRepository) Test() (string, error) {
	dbContext, err := s.CreateUserDbContext()
	if err != nil {
		return "", err
	}
	defer dbContext.Dispose()

	return dbContext.TestQuery()
}