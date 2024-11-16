package service

import (
	"context"
	"dev/master/entity"
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	key, err := s.keyRepository.GetKeyBySubscription(ctx, keyRequest.SubscriptionId)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoKey(*key), nil
}

func (s *KeyService) GetActiveKeysByUser(ctx context.Context, userId *protogen.UserId) (*protogen.CountriesKeys, error) {
	countriesKeys, err := s.keyRepository.FindActiveUsersKeys(ctx, userId.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	keys := make(map[string]*protogen.Key, len(countriesKeys))

	for country, key := range countriesKeys {
		keys[country] = toProtoKey(*key)
	}
	return &protogen.CountriesKeys{Keys: keys}, nil
}

func toProtoKey(key entity.Key) *protogen.Key {
	return &protogen.Key{
		Id:             key.Id,
		Data:           key.Data,
		KeyType:        protogen.KeyType(key.KeyType),
		SubscriptionId: key.SubscriptionId,
		ProxyId:        key.ProxyId,
	}
}
