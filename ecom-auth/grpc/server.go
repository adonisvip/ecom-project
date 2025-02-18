package grpc

import (
	"context"

	pb "ecom-auth/proto/auth"
  "ecom-auth/handlers"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

// Login xử lý yêu cầu đăng nhập
func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := handlers.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		return &pb.LoginResponse{Message: "Invalid credentials"}, nil
	}
	return &pb.LoginResponse{Token: token, Message: "Login successful"}, nil
}

// ValidateToken kiểm tra JWT
func (s *AuthServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	valid := handlers.ValidateJWT(req.Token)
	return &pb.ValidateTokenResponse{Valid: valid}, nil
}
