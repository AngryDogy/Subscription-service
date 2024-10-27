package service

import (
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type ProxyService struct {
	proxyRepository repository.ProxyRepository
	protogen.UnimplementedProxyServiceServer
}

func NewProxyService(repository repository.Repository) *ProxyService {
	return &ProxyService{proxyRepository: repository}
}

func (s *ProxyService) GetProxies(empty *empty.Empty, stream grpc.ServerStreamingServer[protogen.Proxy]) error {
	proxies, err := s.proxyRepository.GetProxies()
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
