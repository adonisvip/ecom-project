package grpc

import (
	"context"
	"log"

	"ecom-auth/handlers"
	pb "ecom-auth/proto/auth"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

// Login xử lý yêu cầu đăng nhập
func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, refreshToken, err := handlers.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		log.Printf("Login failed for user %s: %v", req.Username, err)
		return &pb.LoginResponse{Token: "", RefreshToken: "", Message: "Invalid credentials"}, nil
	}
	if token == "" {
		log.Printf("Login succeeded but token is empty for user %s", req.Username)
		return &pb.LoginResponse{Token: "", RefreshToken: "", Message: "Failed to generate token"}, nil
	}
	return &pb.LoginResponse{Token: token, RefreshToken: refreshToken, Message: "Login successful"}, nil
}

// Register xử lý yêu cầu đăng ký
func (s *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	err := handlers.RegisterUser(req.Username, req.Password, req.Email)
	if err != nil {
		return &pb.RegisterResponse{Message: err.Error()}, nil
	}
	return &pb.RegisterResponse{Message: "Registration successful"}, nil
}

// ValidateToken kiểm tra JWT
func (s *AuthServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	valid := handlers.ValidateJWT(req.Token)
	return &pb.ValidateTokenResponse{Valid: valid}, nil
}

// RefreshToken tạo access token mới và refresh token mới từ refresh token cũ
func (s *AuthServer) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	newAccessToken, newRefreshToken, err := handlers.RefreshToken(req.RefreshToken)
	if err != nil {
		log.Printf("Refresh token failed: %v", err)
		return &pb.RefreshTokenResponse{Token: "", RefreshToken: "", Message: "Invalid or expired refresh token"}, nil
	}
	return &pb.RefreshTokenResponse{Token: newAccessToken, RefreshToken: newRefreshToken, Message: "Token refreshed successfully"}, nil
}
