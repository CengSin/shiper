#原始的grpc编译命令
#protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/cengsin/shiper/consignment-cli \
#	  proto/consignment/consignment.proto
build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/cengsin/shiper/consignment-cli \
		proto/consignment/consignment.proto
	# 为了使用mdns，需要让客户端也容器化
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .

run:
	docker run --name=consignment-cli -e MICRO_REGISTRY=mdns consignment-cli

clean:
	docker stop consignment-cli
	docker rm consignment-cli
	docker rmi consignment-cli