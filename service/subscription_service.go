package service

import (
	"context"
	"dev/master/entity"
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
	"dev/master/service/proxy"

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
		return nil, err
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
		return nil, err
	}

	proxy, err := s.proxyRepository.GetRandomProxyByCountry(ctx, request.CountryId)
	if err != nil {
		return nil, err
	}

	key, err := s.proxyClient.CreateKey(proxy.Address)
	if err != nil {
		return nil, err
	}
	key.SubscriptionId = subscription.Id
	key.ProxyId = proxy.Id

	_, err = s.keyRepository.InsertKey(ctx, *key)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	proxy, err := s.proxyRepository.GetProxyById(ctx, key.ProxyId)
	if err != nil {
		return nil, err
	}

	_ = s.proxyClient.DeleteKey(proxy.Address, key.IdInProxy)

	return nil, nil
}

func (s *SubscriptionService) GetAllSubscriptions(context.Context, *emptypb.Empty) (*protogen.Subscriptions, error) {
	allSubscriptions, err := s.subscriptionRepository.GetAllSubscriptions(context.Background())
	if err != nil {
		return nil, err
	}

	subscriptions := make([]*protogen.Subscription, 0, len(allSubscriptions))
	for _, sub := range allSubscriptions {
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
