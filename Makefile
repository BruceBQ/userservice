GO=go
PLATFORM_FILES="./cmd/main.go"
run-server: 
	@echo Running Userservice for development
	$(GO) run $(PLATFORM_FILES)

build:
	$(GO) build -o userservice $(PLATFORM_FILES)

docker:
	docker build -t userservice .

generate-pb:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/parking.proto