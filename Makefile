
BASEDIR=$(PWD)

all: gen-code build

build:
	go build

gen-code:
	@echo "Generating Go code..."
	@protoc -I svc/ svc/authkm.proto --go_out=plugins=grpc:grpc
	@protoc -I svc/ svc/contentkm.proto --go_out=plugins=grpc:grpc
