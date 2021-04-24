package repository

import (
	"github.com/ujinjinjin/services/user/factories"
	pb "github.com/ujinjinjin/services/user/interface"
	"log"
)

type UserRepository struct {
	factory *factories.DbContextFactory
}

// NewUserRepository create new user repository
func NewUserRepository(factory *factories.DbContextFactory) *UserRepository {
	return &UserRepository{
		factory,
	}
}

// GetUser get user
func (s *UserRepository) GetUser(userId int32) (pb.User, error) {
	dbContext := s.factory.CreateUserDbContext()
	defer dbContext.Dispose()

	dbResult, err := dbContext.GetUser(userId)

	if err != nil {
		log.Printf("UserRepository:GetUser:Error; %v\n", err)
		return pb.User{}, err
	}

	return pb.User{
		UserId: dbResult.UserId,
		FirstName: dbResult.FirstName,
		LastName: dbResult.LastName,
		MiddleName: dbResult.MiddleName,
		Email: dbResult.Email,
	}, nil
}
