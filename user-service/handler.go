package main

import (
	pb "github.com/cengsin/shiper/user-service/proto/user"
	"golang.org/x/crypto/bcrypt"
	context "golang.org/x/net/context"
	"log"
)

type service struct {
	repo         Repository
	tokenService Authable
}

func (srv *service) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {
	if err := srv.repo.Create(req); err != nil {
		return err
	}
	resp.User = req
	return nil
}

func (srv *service) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	resp.User = user
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}

	resp.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {
	log.Println("Logging in with: ", req.Email, req.Password)
	user, err := srv.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}

	//Compare our given password against the hashed password
	//stored in the database
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}
	resp.Token = token
	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	return nil
}
