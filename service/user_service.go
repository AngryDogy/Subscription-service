package service

import (
	"context"
	"dev/master/entity"
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

func (s *UserService) RegisterUser(ctx context.Context, userCreateRequest *protogen.UserCreateRequest) (*protogen.User, error) {
	user, err := s.userRepository.CreateUser(ctx, &entity.User{
		Id:       userCreateRequest.Id,
		Username: userCreateRequest.Username,
	})
	if err != nil {
		return nil, nil
	}

	return &protogen.User{
		Id:       user.Id,
		Username: user.Username,
		HadTrial: user.HadTrial,
	}, nil
}

func (s *UserService) GetUserByID(ctx context.Context, userId *protogen.UserId) (*protogen.User, error) {
	user, err := s.userRepository.FindUserById(ctx, userId.GetId())
	if err != nil {
		return nil, nil
	}

	return &protogen.User{
		Id:       user.Id,
		Username: user.Username,
		HadTrial: user.HadTrial,
	}, nil
}
