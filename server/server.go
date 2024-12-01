package server

import (
	"context"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	logger     *zap.Logger
	grpcServer *grpc.Server
}

func NewServer(logger *zap.Logger, opt []grpc.ServerOption) *Server {
	server := &Server{
		logger: logger,
	}

	opt = append(opt, grpc.UnaryInterceptor(server.logInterception))
	server.grpcServer = grpc.NewServer(opt...)

	return server
}

func (s *Server) SetService(serviceDesc *grpc.ServiceDesc, service interface{}) *Server {
	s.logger.Info("new service was registered", zap.String("service name", serviceDesc.ServiceName))
	s.grpcServer.RegisterService(serviceDesc, service)
	return s
}

func (s *Server) Serve(lis net.Listener) error {
	s.logger.Info("server started", zap.String("address", lis.Addr().String()))
	return s.grpcServer.Serve(lis)
}

func (s *Server) Stop() {
	s.grpcServer.Stop()
}

func (s *Server) logInterception(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	s.logger.Info("gRPC method was called", zap.Any("method", info.FullMethod))
	resp, err := handler(ctx, req)
	if err != nil {
		s.logger.Error("method %q failed: %s", zap.Any("method", info.FullMethod), zap.Error(err))
	}
	return resp, err
}
