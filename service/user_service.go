package service

import (
	"context"
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	userRepository repository.UserRepository
	protogen.UnimplementedUserServiceServer
}

func NewUserService(repository repository.Repository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, userId *protogen.UserId) (*protogen.User, error) {
	user, err := s.userRepository.CreateUser(ctx, userId.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protogen.User{
		Id:       user.Id,
		HadTrial: user.HadTrial,
	}, nil
}

func (s *UserService) GetUserByID(ctx context.Context, userId *protogen.UserId) (*protogen.User, error) {
	user, err := s.userRepository.FindUserById(ctx, userId.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &protogen.User{
		Id:       user.Id,
		HadTrial: user.HadTrial,
	}, nil
}
