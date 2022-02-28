package main

import (
	pb "github.com/cengsin/shiper/user-service/proto/user"
	"github.com/micro/go-micro"
	"log"
)

const schema = `
	create table if not exists users (
		id varchar(36) not null,
		name varchar(125) not null,
		email varchar(225) not null unique,
		password varchar(225) not null,
		company varchar(125),
		primary key (id)
	);
`

func main() {

	// Creates a database connection and handles
	// closing it again before exit.
	db, err := NewConnection()
	if err != nil {
		log.Panic(err)
	}

	session, _ := db.DB()
	defer session.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.Raw(schema)

	repo := NewPostgresRepository(db)

	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"), micro.Version("latest"))

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})

	if err := srv.Run(); err != nil {
		log.Panic(err)
	}
}
