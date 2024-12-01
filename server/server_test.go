package server

import (
	"context"
	"dev/master/entity"
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
	"dev/master/service"
	"dev/master/service/proxy"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const testServerPort = "localhost:8080"

type Client struct {
	keyClient          protogen.KeyServiceClient
	countryClient      protogen.CountryServiceClient
	proxyClient        protogen.ProxyServiceClient
	subscriptionClient protogen.SubscriptionServiceClient
	userClient         protogen.UserServiceClient
}

func initServer(t *testing.T) (*Client, *repository.MockRepository, *proxy.MockClient, func()) {
	logger := zap.Must(zap.NewDevelopment())

	mockCtl := gomock.NewController(t)

	mockRepositry := repository.NewMockRepository(mockCtl)
	mockProxy := proxy.NewMockClient(mockCtl)

	grpcServer := NewServer(logger, make([]grpc.ServerOption, 0))
	grpcServer.
		SetService(&protogen.UserService_ServiceDesc, service.NewUserService(mockRepositry)).
		SetService(&protogen.KeyService_ServiceDesc, service.NewKeyService(mockRepositry)).
		SetService(&protogen.CountryService_ServiceDesc, service.NewCountryService(mockRepositry)).
		SetService(&protogen.ProxyService_ServiceDesc, service.NewProxyService(mockRepositry)).
		SetService(&protogen.SubscriptionService_ServiceDesc, service.NewSubscriptionService(mockRepositry, mockProxy))

	lis, err := net.Listen("tcp", testServerPort)
	require.NoError(t, err)

	go func() {
		grpcServer.Serve(lis)
	}()

	conn, err := grpc.NewClient(testServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	client := Client{
		keyClient:          protogen.NewKeyServiceClient(conn),
		countryClient:      protogen.NewCountryServiceClient(conn),
		proxyClient:        protogen.NewProxyServiceClient(conn),
		subscriptionClient: protogen.NewSubscriptionServiceClient(conn),
		userClient:         protogen.NewUserServiceClient(conn),
	}

	return &client, mockRepositry, mockProxy, func() {
		mockCtl.Finish()
		conn.Close()
		grpcServer.Stop()
	}
}

func TestRegisterUserAndActivateSubscription(t *testing.T) {
	client, repository, proxy, stop := initServer(t)
	defer stop()

	repository.EXPECT().CreateUser(gomock.Any(), int64(1266306390)).Return(&entity.User{}, nil)
	client.userClient.RegisterUser(context.Background(), &protogen.UserId{
		Id: 1266306390,
	})

	expirationDate := time.Now().Add(time.Hour * 168).UTC()
	repository.EXPECT().CreateSubscription(gomock.Any(), entity.Subscription{
		UserId:             1266306390,
		CountryId:          0,
		ExpirationDateTime: expirationDate}).Return(&entity.Subscription{}, nil)

	key := entity.Key{
		ProxyId: 1,
	}
	repository.EXPECT().InsertKey(gomock.Any(), key).Return(&key, nil)
	proxy.EXPECT().CreateKey(gomock.Any()).Return(&key, nil)

	client.subscriptionClient.ActivateSubscription(context.Background(), &protogen.CreateSubscriptionRequest{
		UserId:             1266306390,
		CountryId:          0,
		Trial:              false,
		ExpirationDatetime: timestamppb.New(expirationDate),
	})

	repository.EXPECT().GetKeyBySubscription(gomock.Any(), int64(1)).Return(&entity.Key{}, nil)
	client.keyClient.GetKey(context.Background(), &protogen.KeyRequest{
		SubscriptionId: 1,
	})
}
