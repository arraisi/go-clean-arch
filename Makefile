export NOW=$(shell date +"%Y/%m/%d %T")
export REPO_NAME=astro-erp

build-grpc:
	@echo "${NOW} Building GRPC Server"
	@go build -o ./bin/${REPO_NAME}-grpc cmd/grpc/main.go

run-grpc: build-grpc
	@./bin/${REPO_NAME}-grpc

new-migration:
	@migrate create -ext sql -dir script/migration -seq $(name)

migration-up:
	@migrate -database "${POSTGRES_MASTER_ADDRESS}" -path script/migration up $(step)

migration-down:
	@migrate -database "${POSTGRES_MASTER_ADDRESS}" -path script/migration down $(step)

migration-force:
	@migrate -database "${POSTGRES_MASTER_ADDRESS}" -path script/migration force $(version)