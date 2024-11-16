package service

import (
	"context"
	"dev/master/entity"
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SubscriptionService struct {
	subscriptionRepository repository.SubscriptionRepository
	protogen.UnimplementedSubscriptionServiceServer
}

func NewSubscriptionService(repository repository.Repository) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepository: repository,
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
		})
	}
	return &protogen.Subscriptions{Subscriptions: subscriptions}, nil
}

func (s *SubscriptionService) ActivateSubscription(ctx context.Context, request *protogen.CreateSubscriptionRequest) (*protogen.Subscription, error) {
	subscription, err := s.subscriptionRepository.CreateSubscription(ctx, entity.Subscription{
		UserId:             request.UserId,
		CountryId:          request.CountryId,
		ExpirationDateTime: request.ExpirationDatetime.AsTime(),
	})
	if err != nil {
		return nil, err
	}

	return &protogen.Subscription{
		Id:                 subscription.Id,
		UserId:             subscription.UserId,
		CountryId:          subscription.CountryId,
		ExpirationDatetime: timestamppb.New(subscription.ExpirationDateTime),
	}, nil
}
