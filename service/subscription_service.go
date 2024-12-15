package service

import (
	"context"
	"dev/master/entity"
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
	"dev/master/service/proxy"
	"os"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SubscriptionService struct {
	subscriptionRepository repository.SubscriptionRepository
	keyRepository          repository.KeyRepository
	proxyRepository        repository.ProxyRepository
	proxyClient            proxy.Client
	protogen.UnimplementedSubscriptionServiceServer
}

func NewSubscriptionService(repository repository.Repository, proxyClient proxy.Client) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepository: repository,
		keyRepository:          repository,
		proxyRepository:        repository,
		proxyClient:            proxyClient,
	}
}

func (s *SubscriptionService) GetSubscriptions(ctx context.Context, request *protogen.GetSubscriptionsRequest) (*protogen.Subscriptions, error) {
	subscription, err := s.subscriptionRepository.FindSubscriptions(ctx, request.UserId, request.CountryId, request.Active)
	if err != nil {
		return nil, nil
	}

	subscriptions := make([]*protogen.Subscription, 0, len(subscription))
	for _, sub := range subscription {
		subscriptions = append(subscriptions, &protogen.Subscription{
			Id:                 sub.Id,
			UserId:             sub.UserId,
			CountryId:          sub.CountryId,
			ExpirationDatetime: timestamppb.New(sub.ExpirationDateTime),
			Trial:              sub.IsTrial,
		})
	}
	return &protogen.Subscriptions{Subscriptions: subscriptions}, nil
}

func (s *SubscriptionService) ActivateSubscription(ctx context.Context, request *protogen.CreateSubscriptionRequest) (*protogen.Subscription, error) {
	var subscription *entity.Subscription
	var err error
	if request.Trial {
		subscription, err = s.subscriptionRepository.CreateTrialSubscription(ctx, entity.Subscription{
			UserId:             request.UserId,
			CountryId:          request.CountryId,
			ExpirationDateTime: request.ExpirationDatetime.AsTime(),
			IsTrial:            true,
		})
	} else {
		subscription, err = s.subscriptionRepository.CreateSubscription(ctx, entity.Subscription{
			UserId:             request.UserId,
			CountryId:          request.CountryId,
			ExpirationDateTime: request.ExpirationDatetime.AsTime(),
			IsTrial:            false,
		})
	}

	if err != nil {
		return nil, nil
	}

	proxy, err := s.proxyRepository.GetRandomProxyByCountry(ctx, request.CountryId)
	if err != nil {
		return nil, nil
	}

	key, err := s.proxyClient.CreateKey(proxy.Address)
	if err != nil {
		return nil, nil
	}
	key.SubscriptionId = subscription.Id
	key.ProxyId = proxy.Id

	_, err = s.keyRepository.InsertKey(ctx, *key)
	if err != nil {
		return nil, nil
	}

	return &protogen.Subscription{
		Id:                 subscription.Id,
		UserId:             subscription.UserId,
		CountryId:          subscription.CountryId,
		ExpirationDatetime: timestamppb.New(subscription.ExpirationDateTime),
		Trial:              subscription.IsTrial,
	}, nil
}

func (s *SubscriptionService) DeactivateSubscription(ctx context.Context, request *protogen.DeactivateSubscriptionRequest) (*emptypb.Empty, error) {
	key, err := s.keyRepository.GetKeyBySubscription(ctx, request.SubscriptionId)

	if err != nil {
		return nil, nil
	}

	_ = s.proxyClient.DeleteKey(os.Getenv("DEFAULT_PROXY_ADDRESS"), key.IdInProxy)

	return nil, nil
}
