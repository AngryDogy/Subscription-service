package service

import (
	"context"
	"dev/master/entity"
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
	"dev/master/service/proxy"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
)

const DEFAULT_PROXY_ID = int64(1)

type SubscriptionService struct {
	subscriptionRepository repository.SubscriptionRepository
	keyRepository          repository.KeyRepository
	proxyClient            proxy.Client
	protogen.UnimplementedSubscriptionServiceServer
}

func NewSubscriptionService(repository repository.Repository, proxyClient proxy.Client) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepository: repository,
		keyRepository:          repository,
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
		})
	} else {
		subscription, err = s.subscriptionRepository.CreateSubscription(ctx, entity.Subscription{
			UserId:             request.UserId,
			CountryId:          request.CountryId,
			ExpirationDateTime: request.ExpirationDatetime.AsTime(),
		})
	}

	if err != nil {
		return nil, nil
	}

	key, err := s.proxyClient.CreateKey(os.Getenv("DEFAULT_PROXY_ADDRESS"))
	if err != nil {
		return nil, nil
	}
	key.SubscriptionId = subscription.Id
	key.ProxyId = DEFAULT_PROXY_ID

	_, err = s.keyRepository.InsertKey(ctx, *key)
	if err != nil {
		return nil, nil
	}

	return &protogen.Subscription{
		Id:                 subscription.Id,
		UserId:             subscription.UserId,
		CountryId:          subscription.CountryId,
		ExpirationDatetime: timestamppb.New(subscription.ExpirationDateTime),
	}, nil
}
