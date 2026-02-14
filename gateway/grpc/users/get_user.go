package grpc_users

import (
	"context"
	userspb "diploma/proto/users"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func GRPC_Get_All_Users() ([]*userspb.User, error) {
	conn, err := grpc.Dial("users:50053", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := userspb.NewUsersServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.GetAllUsers(ctx, &userspb.GetAllUsersRequest{})
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, fmt.Errorf("empty grpc response")
	}

	return resp.Users, nil
}
