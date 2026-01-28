package main

import (
	"context"
	authpb "diploma/proto/auth"
	"log"
	"net"

	"google.golang.org/grpc"
)

type AuthService struct {
	authpb.UnimplementedAuthServiceServer
}

func (s *AuthService) SignIn(ctx context.Context, req *authpb.SignInRequest) (*authpb.SignInResponse, error) {
	return &authpb.SignInResponse{Status: true, Token: "Hello, world!"}, nil
}

func main() {
	ls, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Error at gRPC Auth server.", err)
	}

	grpcServer := grpc.NewServer()
	authpb.RegisterAuthServiceServer(grpcServer, &AuthService{})

	log.Println("gRPC Auth server is listening on port :50051")
	if err := grpcServer.Serve(ls); err != nil {
		log.Fatal("Couldn't start gRPC auth server.", err)
	}
}
