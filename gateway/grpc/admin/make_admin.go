package grpc_admin

import (
	"context"
	adminpb "diploma/proto/admin"
	"log"
	"time"

	"google.golang.org/grpc"
)

func GRPC_Make_Admin(user_id int) (bool, string) {
	conn, err := grpc.Dial("admin:50052", grpc.WithInsecure())
	if err != nil {
		log.Println("Couldn't connect to gRPC admin-service server", err)
		return false, "Bad request"
	}
	defer conn.Close()

	client := adminpb.NewAdminServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.MakeAdmin(ctx, &adminpb.MakeAdminRequest{
		Id: int32(user_id),
	})

	if err != nil {
		log.Println("Bad request to admin-service server", err)
		return false, "Bad request"
	}

	return resp.Status, resp.Response
}
