package grpc_auth

import (
	"context"
	authpb "diploma/proto/auth"
	"log"
	"time"

	"google.golang.org/grpc"
)

func GRPC_SignIn(username, password string) (bool, string) {
	conn, err := grpc.Dial("auth:50051", grpc.WithInsecure())
	if err != nil {
		log.Println("Couldn't connect to gRPC auth server", err)
		return false, ""
	}
	defer conn.Close()

	client := authpb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Login(ctx, &authpb.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Println("Couldn't connect to gRPC server", err)
		return false, ""
	}

	return resp.Status, resp.Token
}
