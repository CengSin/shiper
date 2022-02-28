package main

import (
	"context"
	"encoding/json"
	pb "github.com/cengsin/shiper/consignment-cli/proto/consignment"
	microClient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"io/ioutil"
	"log"
	"os"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bytes, &consignment)
	return consignment, nil
}

func main() {

	cmd.Init()
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microClient.DefaultClient)

	/*dialOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(address, dialOption)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewShippingServiceClient(conn)
	*/
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)
	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
