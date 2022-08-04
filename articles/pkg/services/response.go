package services

import "github.com/maslow123/buddyku-users/pkg/pb"

func genericResponse(statusCode int, errorMessage string) (*pb.GenericResponse, error) {
	return &pb.GenericResponse{
		Status: int64(statusCode),
		Error:  errorMessage,
	}, nil
}

func genericUserArticlePointResponse(statusCode int, errorMessage string) (*pb.UserArticlePointResponse, error) {
	return &pb.UserArticlePointResponse{
		Status: int64(statusCode),
		Error:  errorMessage,
	}, nil
}
