package main

import (
	"context"
	userspb "diploma/proto/users"
	"diploma/users-service/db"
	"log"
	"net"

	"google.golang.org/grpc"
)

type UsersService struct {
	userspb.UnimplementedUsersServiceServer
	DB *db.ClinicDB
}

func Sex(s userspb.Sex) string {
	switch s {
	case userspb.Sex_MALE:
		return "М"
	case userspb.Sex_FEMALE:
		return "Ж"
	default:
		return "UNKNOWN"
	}
}

func (s *UsersService) CreateUser(ctx context.Context, req *userspb.CreateUserRequest) (*userspb.CreateUserResponse, error) {
	sex := Sex(req.Sex)
	userId, err := s.DB.Users.CreateUser(req.Username, req.Password, req.Firstname, req.Lastname, req.Email, sex, req.Age)
	if err != nil {
		return &userspb.CreateUserResponse{Status: false, UserId: 0}, nil
	}

	return &userspb.CreateUserResponse{Status: true, UserId: userId}, nil
}

func (s *UsersService) GetAllUsers(ctx context.Context, req *userspb.GetAllUsersRequest) (*userspb.GetAllUsersResponse, error) {
	users, err := s.DB.Users.GetAllUsers()
	if err != nil {
		return &userspb.GetAllUsersResponse{Users: nil}, err
	}
	return &userspb.GetAllUsersResponse{Users: users}, nil
}

func main() {
	conn, err := db.Conn()
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	ls, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatal("Error starting gRPC UsersService server:", err)
	}

	grpcServer := grpc.NewServer()
	userspb.RegisterUsersServiceServer(grpcServer, &UsersService{
		DB: conn,
	})

	log.Println("gRPC Users server is listening on port :50053")
	if err := grpcServer.Serve(ls); err != nil {
		log.Fatal("Couldn't start gRPC UsersService server:", err)
	}
}
