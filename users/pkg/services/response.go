package services

import "github.com/maslow123/buddyku-users/pkg/pb"

func genericResponse(statusCode int, errorMessage string) (*pb.GenericResponse, error) {
	return &pb.GenericResponse{
		Status: int64(statusCode),
		Error:  errorMessage,
	}, nil
}

func genericLoginResponse(statusCode int, errorMessage string) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		Status: int64(statusCode),
		Error:  errorMessage,
	}, nil
}

func genericShowRegisterResponse(statusCode int, errorMessage string) (*pb.ShowUserResponse, error) {
	return &pb.ShowUserResponse{
		Status: int64(statusCode),
		Error:  errorMessage,
	}, nil
}
