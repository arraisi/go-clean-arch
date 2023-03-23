package main

import (
	"go-clean-arch/server"
	"log"
)

func main() {
	log.Println("starting astro-erp-grpc.. ")
	// init all DI for grpc service handler implementation
	if err := server.InitGrpc(); err != nil {
		log.Fatal(err.Error())
	}
}
