GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
PROTOC=protoc

.PHONY: build
build:
	$(GOBUILD) -o ./server/main ./server/main.go && \
	$(GOBUILD) -o ./client/main ./client/main.go

.PHONY: gen-hasher
gen-hasher:
	$(PROTOC) --go_out=. --go_opt=paths=import \
              --go-grpc_out=. --go-grpc_opt=paths=import \
              proto/hasher.proto