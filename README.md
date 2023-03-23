# Golang Clean Architecture

## Service Overview
In this repository there are several binaries :

| Binary Name     | Main File        | Description            | Dev run command |
|-----------------|------------------|------------------------| --------------- |
| clean-arch-grpc | cmd/grpc/main.go | Expose GRPC server for | make run-grpc   |

## Getting Started

### Prerequisite
1. [Git](https://git-scm.com/)
2. [Go 1.19](https://golang.org/)
3. [Docker](https://www.docker.com/)

### Environment Variable for Kube SecretKey
```
export POSTGRES_ADMIN_MASTER_ADDRESS=postgres://user:root@localhost:8000/demo?sslmode=disable
```

### How to install required local infrastructure
```
$ docker-compose up
```

### Running
1. Clone this repo
2. Export Environment Variable Locally
3. Install all dependency `go mod tidy`
4. Try run Application Restful API with `make run` and access http://localhost:8080

