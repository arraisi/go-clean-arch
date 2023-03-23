export NOW=$(shell date +"%Y/%m/%d %T")
export REPO_NAME=astro-erp

build-grpc:
	@echo "${NOW} Building GRPC Server"
	@go build -o ./bin/${REPO_NAME}-grpc cmd/grpc/main.go

run-grpc: build-grpc
	@./bin/${REPO_NAME}-grpc
