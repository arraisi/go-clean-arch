package grpc

import (
	"context"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	server *grpc.Server
}

func (s *Server) Serve(_ context.Context, port string) error {
	s.server = grpc.NewServer(
		grpc.StreamInterceptor(middleware.ChainStreamServer(
			prometheus.StreamServerInterceptor,
		)),
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			prometheus.UnaryServerInterceptor,
			//controller.ServerInterceptor,
		)))

	reflection.Register(s.server)

	// Initialize all metrics.
	prometheus.EnableHandlingTimeHistogram()
	prometheus.Register(s.server)

	// Create port listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// Start gRPC server
	return s.server.Serve(lis)
}

func (s *Server) GracefulStop() {
	s.server.GracefulStop()
}
