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
func (s *UserRpcService) GetUser(_ context.Context, request *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user, err := s.repository.GetUser(request.UserId, int32(request.IsDeleted))
	if err != nil {
		return nil, err
	}

	var result = &pb.GetUserReply{
		User: &user,
	}
	return result, nil
}

// GetUserList get user list by filter
func (s *UserRpcService) GetUserList(_ context.Context, request *pb.GetUserListRequest) (*pb.GetUserListReply, error) {
	userList, err := s.repository.GetUserList(
		request.UserIdList,
		request.Username,
		request.Email,
		int32(request.IsDeleted),
	)
	if err != nil {
		return nil, err
	}

	var result = &pb.GetUserListReply{
		UserList: userList,
	}
	return result, nil
}
