package services

import (
	"context"
	"log"

	pb "github.com/ujinjinjin/services/user/interface"
	"github.com/ujinjinjin/services/user/repository"
)

type UserRpcService struct {
	pb.UnimplementedUserServiceServer
	repository *repository.UserRepository
}

func InitService(userRepository *repository.UserRepository) *UserRpcService {
	return &UserRpcService{
		repository: userRepository,
	}
}

// CreateUser creates user with specified fields
func (s *UserRpcService) CreateUser(context context.Context, request *pb.CreateUserRequest) (*pb.CreateUserReply, error) {

	log.Print("New user created:")
	log.Printf("\tEmail: %s", request.Email)
	log.Printf("\tFirstName: %s", request.FirstName)
	log.Printf("\tLastName: %s", request.LastName)
	log.Printf("\tMiddleName: %s", request.MiddleName)

	testResult, err := s.repository.Test()
	if err != nil {
		return nil, err
	}

	log.Printf("\tRepository response: %s", testResult)

	var result = &pb.CreateUserReply{
		Id: 1,
	}
	return result, nil
}
