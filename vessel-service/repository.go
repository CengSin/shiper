package main

import (
	"errors"
	pb "github.com/cengsin/shiper/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessel"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.Collection().Insert(vessel)
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessels []*pb.Vessel
	err := repo.Collection().Find(nil).All(&vessels)
	if err != nil {
		return nil, err
	}

	for _, v := range vessels {
		if spec.MaxWeight <= v.MaxWeight && spec.Capacity <= v.Capacity {
			return v, nil
		}
	}

	return nil, errors.New("No vessel found by that spec")
}

func (repo *VesselRepository) Collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}

func (repo *VesselRepository) Close() {
	repo.session.Close()
}

//type Repository interface {
//	FindAvailable(*pb.Specification) (*pb.Vessel, error)
//}
//
//type VesselRepository struct {
//	vessels []*pb.Vessel
//}
//
//func (v *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
//	for _, vessel := range v.vessels {
//		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
//			return vessel, nil
//		}
//	}
//	return nil, errors.New("No vessel found by that spec")
//}
