package auth

import (
	"context"
	authpb "diploma/proto/auth"
	"google.golang.org/grpc"
	"log"
	"time"
)

func SignIn(username, password string) (bool, string) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return false, "Couldn't connect to gRPC server"
	}
	defer conn.Close()

	client := authpb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.SignIn(ctx, &authpb.SignInRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Println(err)
		return false, "Couldn't request 'SignIn' service"
	}

	return resp.Status, resp.Token
}
