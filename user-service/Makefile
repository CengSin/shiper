build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/cengsin/shiper/user-service \
    		proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build -o user-service .
	docker build -t user-service .

run:
	docker run -p 50057:50051 --name=user-service \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns user-service


clean:
	docker stop user-service
	docker rm user-service
	docker rmi user-service