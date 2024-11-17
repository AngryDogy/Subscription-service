package main

import (
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
	"dev/master/server"
	"dev/master/service"
	"dev/master/service/proxy"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()

	//err := godotenv.Load()
	//if err != nil {
	//	logger.Fatal("Error loading .env file", zap.Error(err))
	//}

	postgresRepository := repository.NewPostgresRepository(logger)
	err := postgresRepository.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Fatal("Error connecting to postgres database", zap.Error(err))
	}
	defer postgresRepository.Close()

	//err = postgresRepository.Initialize(os.Getenv("SCHEME_PATH"))
	//if err != nil {
	//	logger.Fatal("Error initializing postgres database", zap.Error(err))
	//}

	grpcServer := server.NewServer(logger, make([]grpc.ServerOption, 0))
	grpcServer.
		SetService(&protogen.UserService_ServiceDesc, service.NewUserService(postgresRepository)).
		SetService(&protogen.KeyService_ServiceDesc, service.NewKeyService(postgresRepository)).
		SetService(&protogen.CountryService_ServiceDesc, service.NewCountryService(postgresRepository)).
		SetService(&protogen.ProxyService_ServiceDesc, service.NewProxyService(postgresRepository)).
		SetService(&protogen.SubscriptionService_ServiceDesc, service.NewSubscriptionService(postgresRepository, proxy.NewClient()))

	lis, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		logger.Fatal("Failed to create listener", zap.Error(err))
	}
	grpcServer.Serve(lis)
}
