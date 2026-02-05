package grpc_auth

import (
	"context"
	userspb "diploma/proto/users"
	"log"
	"time"

	"google.golang.org/grpc"
)

func GRPC_SignUp(username, password, firstname, lastname, email, sex, age string) (bool, int32) {
	conn, err := grpc.Dial("users:50053", grpc.WithInsecure())
	if err != nil {
		log.Println("Couldn't connect to gRPC auth server", err)
		return false, 0
	}
	defer conn.Close()

	client := userspb.NewUsersServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// string to enum like in proto
	var sexEnum userspb.Sex
	switch sex {
	case "М":
		sexEnum = userspb.Sex_MALE
	case "Ж":
		sexEnum = userspb.Sex_FEMALE
	default:
		sexEnum = userspb.Sex_UNKNOWN
	}

	resp, err := client.CreateUser(ctx, &userspb.CreateUserRequest{
		Username:  username,
		Password:  password,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Sex:       sexEnum,
		Age:       age,
	})

	if err != nil {
		log.Println("Couldn't connect to gRPC server", err)
		return false, 0
	}

	return resp.Status, resp.UserId
}
