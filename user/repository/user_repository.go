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
func (s *UserRepository) GetUser(userId int32, isDeleted int32) (pb.User, error) {
	dbContext := s.factory.CreateUserDbContext()
	defer dbContext.Dispose()

	dbResult, err := dbContext.GetUser(userId, isDeleted)
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
		IsDeleted: dbResult.IsDeleted,
	}, nil
}

// GetUserList get user list by filter
func (s *UserRepository) GetUserList(userIdList []int32, username string, email string, isDeleted int32) ([]*pb.User, error) {
	dbContext := s.factory.CreateUserDbContext()
	defer dbContext.Dispose()

	dbResult, err := dbContext.GetUserList(userIdList, username, email, isDeleted)
	if err != nil {
		log.Printf("UserRepository:GetUser:Error; %v\n", err)
		return nil, err
	}

	userList := make([]*pb.User, len(dbResult))
	for i := 0; i < len(userList); i++ {
		var user = &pb.User{
			UserId: dbResult[i].UserId,
			FirstName: dbResult[i].FirstName,
			LastName: dbResult[i].LastName,
			MiddleName: dbResult[i].MiddleName,
			Email: dbResult[i].Email,
			IsDeleted: dbResult[i].IsDeleted,
		}
		userList[i] = user
	}

	return userList, nil
}
