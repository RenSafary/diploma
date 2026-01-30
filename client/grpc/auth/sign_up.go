package grpc_auth

import (
	"context"
	authpb "diploma/proto/auth"
	"google.golang.org/grpc"
	"log"
	"time"
)

func GRPC_SignUp(username, password, firstname, lastname, email, sex, age string) (bool, string) {
	conn, err := grpc.Dial("auth:50051", grpc.WithInsecure())
	if err != nil {
		log.Println("Couldn't connect to gRPC auth server", err)
		return false, ""
	}
	defer conn.Close()

	client := authpb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.SignUp(ctx, &authpb.SignUpRequest{
		Username:  username,
		Password:  password,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Sex:       sex,
		Age:       age,
	})

	if err != nil {
		log.Println("Couldn't connect to gRPC server", err)
		return false, ""
	}

	return resp.Status, resp.Token
}
