package service

import (
	"context"
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
	"google.golang.org/grpc"
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

func (s *SubscriptionService) GetUserSubscriptions(userID *protogen.UserID, stream grpc.ServerStreamingServer[protogen.Subscription]) error {
	subs, err := s.subscriptionRepository.FindSubscriptionsByUserID(userID.GetId())
	if err != nil {
		return err
	}

	for _, sub := range subs {
		stream.Send(&protogen.Subscription{
			Id:             sub.Id,
			UserId:         sub.UserId,
			CountryId:      sub.CountryId,
			ExpirationTime: &sub.ExpirationTime,
		})
	}
	return nil
}

func (s *SubscriptionService) ActivateTrialSubscription(ctx context.Context, subReq *protogen.SubscriptionRequest) (*protogen.Subscription, error) {
	sub, err := s.subscriptionRepository.CreateTrialSubscription(ctx, subReq.UserId, subReq.CountryId)
	if err != nil {
		return nil, err
	}

	return &protogen.Subscription{
		Id:             sub.Id,
		UserId:         sub.UserId,
		CountryId:      sub.CountryId,
		ExpirationTime: &sub.ExpirationTime,
	}, nil

}

func (s *SubscriptionService) ActivatePaidSubscription(ctx context.Context, subReq *protogen.SubscriptionRequest) (*protogen.Subscription, error) {
	sub, err := s.subscriptionRepository.CreateSubscription(ctx, subReq.UserId, subReq.CountryId)
	if err != nil {
		return nil, err
	}

	return &protogen.Subscription{
		Id:             sub.Id,
		UserId:         sub.UserId,
		CountryId:      sub.CountryId,
		ExpirationTime: &sub.ExpirationTime,
	}, err
}
