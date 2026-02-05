package main

import (
	"context"
	"diploma/auth-service/db"
	authpb "diploma/proto/auth"
	"log"
	"net"

	"google.golang.org/grpc"
)

type AuthService struct {
	authpb.UnimplementedAuthServiceServer
	DB *db.ClinicDB
}

func (s *AuthService) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	token, err := s.DB.Users.GiveToken(req.Username, req.Password)
	if err != nil || token == "Wrong login or password" {
		return &authpb.LoginResponse{Status: false, Token: ""}, err
	}

	return &authpb.LoginResponse{Status: true, Token: token}, nil
}

func main() {
	ls, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Error at gRPC Auth server.", err)
	}

	db, err := db.Conn()
	if err != nil {
		log.Fatal("gRPC server error with DB: ", err)
		return
	}

	grpcServer := grpc.NewServer()
	authpb.RegisterAuthServiceServer(grpcServer, &AuthService{
		DB: db,
	})

	log.Println("gRPC Auth server is listening on port :50051")
	if err := grpcServer.Serve(ls); err != nil {
		log.Fatal("Couldn't start gRPC auth server.", err)
	}
}
