package server

import "go-clean-arch/controller/grpc"

func InitGrpc() error {
	server := grpc.Server{}
	return runGrpcServer(server, "8080")
}
