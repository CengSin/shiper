build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/cengsin/shiper/vessel-service \
		proto/vessel/vessel.proto
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vessel-service .
	docker build -t vessel-service .

run:
	docker run -p 50052:50051 \
		-e MICRO_SERVICE_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		--name=vessel-service \
		vessel-service

clean:
	docker stop vessel-service
	docker rm vessel-service
	docker rmi vessel-service