package main

import (
	pb "github.com/cengsin/shiper/consignment-service/proto/consignment"
	vesselProto "github.com/cengsin/shiper/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"log"
	"os"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}
	defer session.Close()

	// 设置grpc服务器
	//lis, err := net.Listen("tcp", port)
	//if err != nil {
	//	log.Fatalln("failed to listen", err)
	//}
	//s := grpc.NewServer()

	// 使用go-micro替换
	srv := micro.NewService(micro.Name("go.micro.srv.consignment"), micro.Version("latest"))

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(), &service{session: session, vesselClient: vesselClient})

	////在grpc服务器上注册reflection
	//reflection.Register(s)
	//if err = s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
