#原始的grpc编译命令
#protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/cengsin/shiper/consignment-cli \
#	  proto/consignment/consignment.proto
build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/cengsin/shiper/consignment-service \
		proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build -o consignment-service .
	docker build -t consignment-service .

run:
	docker run -p 50051:50051 --name=consignment-service \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns consignment-service


clean:
	docker stop consignment-service
	docker rm consignment-service
	docker rmi consignment-service