package main

import (
	"fmt"
	pb "github.com/cengsin/shiper/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"log"
	"os"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	host := os.Getenv("DB_HOST")
	if len(host) == 0 {
		host = defaultHost
	}

	session, err := CreateSession(host)
	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}
	defer session.Clone()

	//vessels := []*pb.Vessel{{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500}}

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()
	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
