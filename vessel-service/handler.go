package main

import (
	"context"
	pb "github.com/cengsin/shiper/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
)

//type service struct {
//	repo Repository
//}
//
//func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
//	available, err := s.repo.FindAvailable(spec)
//	if err != nil {
//		return err
//	}
//
//	resp.Vessel = available
//	return nil
//}

type service struct {
	session *mgo.Session
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	if err := repo.Create(req); err != nil {
		return err
	}
	resp.Vessel = req
	resp.Created = true
	return nil
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	available, err := repo.FindAvailable(spec)
	if err != nil {
		return err
	}

	resp.Vessel = available
	return nil
}
