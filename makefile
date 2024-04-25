default: build run

build:
	@echo "Building service..."
	@GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o service/build/bin service/cmd/main.go
	@chmod +x service/build/bin
	@echo "Done."
	@echo "Building HTTP proxy..."
	@GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o httpProxy/build/bin httpProxy/cmd/main.go
	@chmod +x httpProxy/build/bin
	@echo "Done."
	@echo "Building worker..."
	@GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o worker/build/bin worker/cmd/main.go
	@chmod +x worker/build/bin
	@echo "Done."

run: build
	@docker compose up

clean:
	@docker compose down -v
	@docker image rmi `docker images -q`
	@docker system prune

gen-proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		service/pkg/proto/service.proto

debug:
	@docker run -it --entrypoint=/bin/sh project-simplas-service
