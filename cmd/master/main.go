package main

import (
	"context"
	"dev/master/entity"
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
	"dev/master/server"
	"dev/master/service"
	"dev/master/service/proxy"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert/yaml"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file", zap.Error(err))
	}

	postgresRepository := repository.NewPostgresRepository(logger)
	err = postgresRepository.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Fatal("Error connecting to postgres database", zap.Error(err))
	}
	defer postgresRepository.Close()

	err = postgresRepository.Initialize(os.Getenv("SCHEME_PATH"))
	if err != nil {
		logger.Fatal("Error initializing postgres database", zap.Error(err))
	}

	err = uploadCountries(postgresRepository, logger)
	if err != nil {
		logger.Fatal("failed to upload counties", zap.Error(err))
	}

	err = uploadProxies(postgresRepository, logger)
	if err != nil {
		logger.Fatal("failed to upload proxies", zap.Error(err))
	}

	grpcServer := server.NewServer(logger, make([]grpc.ServerOption, 0))
	grpcServer.
		SetService(&protogen.UserService_ServiceDesc, service.NewUserService(postgresRepository)).
		SetService(&protogen.KeyService_ServiceDesc, service.NewKeyService(postgresRepository)).
		SetService(&protogen.CountryService_ServiceDesc, service.NewCountryService(postgresRepository)).
		SetService(&protogen.ProxyService_ServiceDesc, service.NewProxyService(postgresRepository)).
		SetService(&protogen.SubscriptionService_ServiceDesc, service.NewSubscriptionService(postgresRepository, proxy.NewClient(logger)))

	lis, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		logger.Fatal("Failed to create listener", zap.Error(err))
	}
	grpcServer.Serve(lis)
}

func uploadCountries(repository repository.Repository, logger *zap.Logger) error {
	var countriesCollection entity.CountryCollection

	yamlFile, err := os.ReadFile(os.Getenv("COUNTRIES_LIST"))
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &countriesCollection)
	if err != nil {
		return err
	}

	for _, country := range countriesCollection.Countries {
		_, err = repository.CreateCountry(context.Background(), country)
		if err != nil {
			logger.Warn("Error occured while creating countries", zap.Error(err))
		} else {
			logger.Info("Country was successfully uploaded", zap.String("country name", country.Name))
		}
	}

	return nil
}

func uploadProxies(repository repository.Repository, logger *zap.Logger) error {
	var proxiesCollection entity.ProxyCollection

	yamlFile, err := os.ReadFile(os.Getenv("PROXIES_LIST"))
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &proxiesCollection)
	if err != nil {
		return err
	}

	for _, proxy := range proxiesCollection.Proxies {
		_, err := repository.CreateProxy(context.Background(), proxy)
		if err != nil {
			logger.Warn("Error occured while creating proxies", zap.Error(err))
		} else {
			logger.Info("Proxe was successfully uploaded", zap.String("address", proxy.Address))
		}
	}

	return nil
}
