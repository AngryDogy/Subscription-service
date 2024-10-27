package service

import (
	"context"
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
)

type KeyService struct {
	keyRepository repository.KeyRepository
	protogen.UnimplementedKeyServiceServer
}

func NewKeyService(repository repository.Repository) *KeyService {
	return &KeyService{
		keyRepository: repository,
	}
}

func (s *KeyService) GetKey(ctx context.Context, keyRequest *protogen.KeyRequest) (*protogen.Key, error) {
	key, err := s.keyRepository.FindUserKey(ctx, keyRequest.UserId, keyRequest.CountryId)
	if err != nil {
		return nil, err
	}

	return &protogen.Key{
		Id:             key.Id,
		Data:           key.Data,
		KeyType:        protogen.KeyType(key.KeyType),
		SubscriptionId: key.SubscriptionId,
	}, nil
}
