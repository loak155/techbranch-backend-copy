package router

import (
	"context"

	"github.com/loak155/techbranch-backend/internal/auth/usecase"
	authpb "github.com/loak155/techbranch-backend/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type IAuthGRPCServer interface {
	Signup(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupResponse, error)
	Signin(ctx context.Context, req *authpb.SigninRequest) (*authpb.SigninResponse, error)
	GenerateToken(ctx context.Context, req *authpb.GenerateTokenRequest) (*authpb.GenerateTokenResponse, error)
	ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error)
	RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error)
}

type authGRPCServer struct {
	authpb.UnimplementedAuthServiceServer
	uu usecase.IAuthUsecase
}

func NewAuthGRPCServer(grpcServer *grpc.Server, uu usecase.IAuthUsecase) authpb.AuthServiceServer {
	s := authGRPCServer{uu: uu}
	authpb.RegisterAuthServiceServer(grpcServer, &s)
	reflection.Register(grpcServer)
	return &s
}

func (s *authGRPCServer) Signup(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupResponse, error) {
	res := authpb.SignupResponse{}
	userRes, err := s.uu.Signup(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	// res.User = &userpb.User{
	// 	Id: int32(userRes),
	// }
	res.User = &authpb.User{
		Id: int32(userRes),
	}
	return &res, nil
}

func (s *authGRPCServer) Signin(ctx context.Context, req *authpb.SigninRequest) (*authpb.SigninResponse, error) {
	res := authpb.SigninResponse{}
	token, err := s.uu.Signin(ctx, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	res.Token = token
	return &res, nil
}

func (s *authGRPCServer) GenerateToken(ctx context.Context, req *authpb.GenerateTokenRequest) (*authpb.GenerateTokenResponse, error) {
	res := authpb.GenerateTokenResponse{}
	authRes, err := s.uu.GenerateToken(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}
	res.Token = authRes
	return &res, nil
}

func (s *authGRPCServer) ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	res := authpb.ValidateTokenResponse{}
	authRes, err := s.uu.ValidateToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	res.Valid = authRes
	return &res, nil
}

func (s *authGRPCServer) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {
	res := authpb.RefreshTokenResponse{}
	authRes, err := s.uu.RefreshToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	res.Token = authRes
	return &res, nil
}
