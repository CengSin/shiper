package main

import (
	pb "github.com/cengsin/shiper/consignment-service/proto/consignment"
	"gopkg.in/mgo.v2"
)

const (
	dbName                = "shippy"
	consignmentCollection = "consignments"
)

type Repository interface {
	Create(consignment *pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

type ConsignmentRepository struct {
	session *mgo.Session
}

func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error {
	return repo.collection().Insert(consignment)
}

func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	// Find()通常接受一个询问条件(query)，但我们想要所有的货运任务，所以在这里用nil
	// 然后把找到的所有货运任务通过All()赋值给consignment
	// 另外在mgo中，One可以处理单个结果
	var consignments []*pb.Consignment
	err := repo.collection().Find(nil).All(&consignments)
	return consignments, err
}

func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}

func (repo *ConsignmentRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(consignmentCollection)
}

// 内容抽象并写入repository.go中
//type IRepository interface {
//	Create(consignment *pb.Consignment) (*pb.Consignment, error)
//	GetAll() []*pb.Consignment
//}
//
//// Repository - 模拟一个数据库，我们会在此后使用真正的数据库替代他
//type Repository struct {
//	consignment []*pb.Consignment
//}
//
//func (r *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
//	r.consignment = append(r.consignment, consignment)
//	return consignment, nil
//}
//
//func (r *Repository) GetAll() []*pb.Consignment {
//	return r.consignment
//}
