package client

import (
	"fmt"

	"github.com/loak155/techbranch-backend/internal/pkg/config"
	pb "github.com/loak155/techbranch-backend/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserGRPCClient(conf *config.Config) (pb.UserServiceClient, error) {
	address := fmt.Sprintf("%s:%d", conf.User.Server.Host, conf.User.Server.Port)
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	return pb.NewUserServiceClient(conn), nil
}

// type IUserGRPCClient interface {
// 	CreateUser(req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
// 	GetUser(req *pb.GetUserRequest) (*pb.GetUserResponse, error)
// 	GetUserByEmail(req *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponse, error)
// }

// type userGRPCClient struct {
// 	client pb.UserServiceClient
// }

// func NewUserGRPCClient() (IUserGRPCClient, error) {
// 	address := fmt.Sprintf("%s:%s", os.Getenv("USER_SERVICE_HOST"), os.Getenv("USER_SERVICE_PORT"))
// 	conn, err := grpc.Dial(
// 		address,
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 		grpc.WithBlock(),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	client := pb.NewUserServiceClient(conn)

// 	return &userGRPCClient{client}, nil
// }

// func (c *userGRPCClient) CreateUser(req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
// 	res, err := c.client.CreateUser(context.Background(), req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return res, nil
// }

// func (c *userGRPCClient) GetUser(req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
// 	res, err := c.client.GetUser(context.Background(), req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return res, nil
// }

// func (c *userGRPCClient) GetUserByEmail(req *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponse, error) {
// 	res, err := c.client.GetUserByEmail(context.Background(), req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return res, nil
// }
