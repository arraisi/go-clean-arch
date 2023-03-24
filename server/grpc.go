package server

import (
	"context"
	"go-clean-arch/controller/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func runGrpcServer(server grpc.Server, port string) error {
	idleConnections := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		signals := make(chan os.Signal, 1)

		// When using socket master, it send SIGTERM after spawning new process,
		// SIGHUP is for handling upstart reload
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
		<-signals

		// We received an os signal, shut down.
		log.Println("GRPC server shutdown gracefully")
		cancel()
		server.GracefulStop()
		close(idleConnections)
	}()

	log.Println("GRPC server running on port", port)
	if err := server.Serve(ctx, port); err != http.ErrServerClosed {
		// Error starting or closing listener:
		return err
	}

	<-idleConnections
	log.Println("GRPC server stopping")
	return nil
}
