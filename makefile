.PHONY: build
build: gen
	go build -o octopus ./cmd/main.go

.PHONY: gen
gen: gen-grpc gen-go

# grpc
.PHONY: gen-grpc
gen-grpc:
	rm -rf ./pkg/grpc/proto
	mkdir -p ./pkg/grpc/proto
	protoc \
    	-I=./proto \
    	--proto_path=./proto \
    	--go_opt=paths=source_relative \
    	--go_out=./pkg/grpc/proto \
    	--go-grpc_opt=paths=source_relative \
    	--go-grpc_out=./pkg/grpc/proto \
    	$(shell find ./proto -iname "*.proto") 2>&1 > /dev/null

.PHONY: gen-go
gen-go:
	go get -u ./...
	go mod tidy
	go generate ./...
	go test -v ./...
	