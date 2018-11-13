
BASEDIR=$(PWD)

all: gen-code build

build:
	go build

gen-code: submodules-update
	@echo "Generating Go code..."
	@protoc -I svc/ svc/authkm.proto --go_out=plugins=grpc:grpc
	@protoc -I svc/ svc/contentkm.proto --go_out=plugins=grpc:grpc
#	@protoc -I svc/ svc/authvanille.proto --go_out=auth
#	@protoc -I svc/ svc/contentvanille.proto --go_out=content
	@echo "done :-)"

submodules-update:
	@sh cmd.sh sub.update
