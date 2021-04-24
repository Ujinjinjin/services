package services

import (
	"context"
	pb "github.com/ujinjinjin/services/user/interface"
	"github.com/ujinjinjin/services/user/repository"
)

type UserRpcService struct {
	pb.UnimplementedUserServiceServer
	repository *repository.UserRepository
}

// InitService initialize rpc service
func InitService(userRepository *repository.UserRepository) *UserRpcService {
	return &UserRpcService{
		repository: userRepository,
	}
}

// GetUser get user
func (s *UserRpcService) GetUser(context context.Context, request *pb.GetUserRequest) (*pb.GetUserReply, error) {

	user, err := s.repository.GetUser(request.UserId)
	if err != nil {
		return nil, err
	}

	var result = &pb.GetUserReply{
		User: &pb.User{
			UserId:     user.UserId,
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			MiddleName: user.MiddleName,
			Email:      user.Email,
		},
	}
	return result, nil
}
