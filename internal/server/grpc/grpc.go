package grpc

import (
	"google.golang.org/grpc"

	"eight/internal/grpc/service"
)

// Grpc struct holds all the dependencies required for starting Grpc server
type Grpc struct {
	conn              *grpc.ClientConn
	UserServiceClient service.UserServiceClient
}

func New() *Grpc {
	serverAddress := "localhost:7080"

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	//defer conn.Close()

	userClient := service.NewUserServiceClient(conn)

	return &Grpc{
		conn:              conn,
		UserServiceClient: userClient,
	}
}
