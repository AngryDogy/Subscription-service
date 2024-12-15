package service

import (
	"context"
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProxyService struct {
	proxyRepository repository.ProxyRepository
	protogen.UnimplementedProxyServiceServer
}

func NewProxyService(repository repository.Repository) *ProxyService {
	return &ProxyService{proxyRepository: repository}
}

func (s *ProxyService) GetProxies(empty *emptypb.Empty, stream grpc.ServerStreamingServer[protogen.Proxy]) error {
	proxies, err := s.proxyRepository.GetAllProxies(context.Background())
	if err != nil {
		return err
	}

	for _, proxy := range proxies {
		stream.Send(&protogen.Proxy{
			Id:        proxy.Id,
			Address:   proxy.Address,
			CountryId: proxy.CountryId,
		})
	}
	return nil
}
