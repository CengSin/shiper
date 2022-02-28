package main

import (
	"context"
	pb "github.com/cengsin/shiper/consignment-service/proto/consignment"
	vesselProto "github.com/cengsin/shiper/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
	"log"
)

type service struct {
	session *mgo.Session
	//repo Repository
	//添加货船服务客户端对象，用来检查是否有符合条件的货船
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{session: s.session.Clone()}
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			log.Printf(e.Error())
		}
	}()
	// 通过货船服务客户端 发出一个 查找是否有合适的货船 的请求
	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("==========%v=======\n", vesselResponse)
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	if err = repo.Create(req); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	getAll, err := repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = getAll
	return nil
}

//使用了go-micro之后这些方法会被替换掉
//func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
//	consignment, err := s.repo.Create(req)
//	if err != nil {
//		return nil, err
//	}
//
//	return &pb.Response{Created: true, Consignment: consignment}, nil
//}
//
//func (s *service) GetConsignments(ctx context.Context, request *pb.GetRequest) (*pb.Response, error) {
//	all := s.repo.GetAll()
//	return &pb.Response{Consignments: all}, nil
//}
