package service

import (
	"context"
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
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

func (s *UserService) RegisterUser(ctx context.Context, userID *protogen.UserID) (*protogen.User, error) {
	user, err := s.userRepository.CreateUser(ctx, userID.GetId())
	if err != nil {
		return nil, err
	}

	return &protogen.User{
		Id:       user.Id,
		HadTrial: user.HadTrial,
	}, nil
}

func (s *UserService) GetUserByID(ctx context.Context, userID *protogen.UserID) (*protogen.User, error) {
	user, err := s.userRepository.FindUserById(ctx, userID.GetId())
	if err != nil {
		return nil, err
	}

	return &protogen.User{
		Id:       user.Id,
		HadTrial: user.HadTrial,
	}, nil
}
