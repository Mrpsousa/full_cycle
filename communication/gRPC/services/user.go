package services

import (
	"context"

	"github.com/mrpsousa/full_cycle/communication/gRPC/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) addUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	return &pb.User{
		Id:    "1",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}
