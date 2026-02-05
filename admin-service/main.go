package main

import (
	"context"
	"diploma/admin-service/db"
	adminpb "diploma/proto/admin"
	"log"
	"net"

	"google.golang.org/grpc"
)

type AdminService struct {
	adminpb.UnimplementedAdminServiceServer
	DB *db.ClinicDB
}

func (s *AdminService) MakeAdmin(ctx context.Context, req *adminpb.MakeAdminRequest) (*adminpb.MakeAdminResponse, error) {
	err := s.DB.Users.MakeAdminDB(int(req.Id))
	if err != nil {
		log.Println(err)
		return &adminpb.MakeAdminResponse{
			Status:   false,
			Response: "Internal Server Error",
		}, err
	}
	return &adminpb.MakeAdminResponse{
		Status:   true,
		Response: "OK",
	}, err
}

func (s *AdminService) DeleteAdmin(ctx context.Context, req *adminpb.DeleteAdminRequest) (*adminpb.DeleteAdminResponse, error) {
	err := s.DB.Users.DeleteAdminDB(int(req.Id))
	if err != nil {
		log.Println(err)
		return &adminpb.DeleteAdminResponse{
			Status:   false,
			Response: "Internal Server Error",
		}, err
	}
	return &adminpb.DeleteAdminResponse{
		Status:   true,
		Response: "OK",
	}, err
}

func main() {
	ls, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal("Error at gRPC Admin service server.", err)
	}

	db, err := db.Conn()
	if err != nil {
		log.Fatal("gRPC server error with DB: ", err)
		return
	}

	grpcServer := grpc.NewServer()
	adminpb.RegisterAdminServiceServer(grpcServer, &AdminService{
		DB: db,
	})

	log.Println("gRPC Admin service server is listening on port :50052")
	if err := grpcServer.Serve(ls); err != nil {
		log.Fatal("Couldn't start gRPC admin service server.", err)
	}
}
