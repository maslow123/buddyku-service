package services

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/maslow123/buddyku-users/pkg/pb"
	"github.com/maslow123/buddyku-users/pkg/utils"
	"github.com/stretchr/testify/require"
)

var randUser, randPass string

func TestRegister(t *testing.T) {
	randUser = utils.RandomString(10)
	randPass = utils.RandomString(10)
	level := int32(0)

	testCases := []struct {
		name string
		req  *pb.RegisterRequest
		resp *pb.GenericResponse
	}{
		{
			"OK -> Register Publisher",
			&pb.RegisterRequest{
				Name:     utils.RandomString(10),
				Username: randUser,
				Password: randPass,
				Level:    level,
			},
			&pb.GenericResponse{
				Status: int64(http.StatusCreated),
				Error:  "",
			},
		},
		{
			"OK -> Register Company",
			&pb.RegisterRequest{
				Name:     utils.RandomString(10),
				Username: fmt.Sprintf("Company %s", randUser),
				Password: randPass,
				Level:    int32(1),
			},
			&pb.GenericResponse{
				Status: int64(http.StatusCreated),
				Error:  "",
			},
		},
		{
			"Invalid Name",
			&pb.RegisterRequest{
				Name:     "",
				Username: randUser,
				Password: randPass,
				Level:    level,
			},
			&pb.GenericResponse{
				Status: int64(http.StatusBadRequest),
				Error:  "invalid-name",
			},
		},
		{
			"Invalid Username",
			&pb.RegisterRequest{
				Name:     utils.RandomString(10),
				Username: "",
				Password: randPass,
				Level:    level,
			},
			&pb.GenericResponse{
				Status: int64(http.StatusBadRequest),
				Error:  "invalid-username",
			},
		},
		{
			"Invalid Password",
			&pb.RegisterRequest{
				Name:     utils.RandomString(10),
				Username: randUser,
				Password: "",
				Level:    level,
			},
			&pb.GenericResponse{
				Status: int64(http.StatusBadRequest),
				Error:  "invalid-password",
			},
		},
		{
			"Invalid Level",
			&pb.RegisterRequest{
				Name:     utils.RandomString(10),
				Username: randUser,
				Password: randPass,
				Level:    int32(2),
			},
			&pb.GenericResponse{
				Status: int64(http.StatusBadRequest),
				Error:  "invalid-level",
			},
		},
		{
			"Existing User",
			&pb.RegisterRequest{
				Name:     utils.RandomString(10),
				Username: randUser,
				Password: randPass,
			},
			&pb.GenericResponse{
				Status: int64(http.StatusBadRequest),
				Error:  "username-already-exists",
			},
		},
	}

	ctx := context.Background()
	conn := checkConnection(ctx, t)
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			response, err := client.Register(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.resp.Status, response.Status)
			require.Equal(t, tc.resp.Error, response.Error)
		})
	}
}

func TestLogin(t *testing.T) {
	randUser = "user1@gmail.com"
	randPass = "111111"

	testCases := []struct {
		name string
		req  *pb.LoginRequest
		resp *pb.LoginResponse
	}{
		{
			"OK",
			&pb.LoginRequest{
				Username: randUser,
				Password: randPass,
			},
			&pb.LoginResponse{
				Status: int64(http.StatusOK),
				Error:  "",
			},
		},
		{
			"Invalid Username",
			&pb.LoginRequest{
				Username: "",
				Password: randPass,
			},
			&pb.LoginResponse{
				Status: int64(http.StatusBadRequest),
				Error:  "invalid-username",
			},
		},
		{
			"Invalid Password",
			&pb.LoginRequest{
				Username: randUser,
				Password: "",
			},
			&pb.LoginResponse{
				Status: int64(http.StatusBadRequest),
				Error:  "invalid-password",
			},
		},
		{
			"Wrong Password",
			&pb.LoginRequest{
				Username: randUser,
				Password: "wrong password",
			},
			&pb.LoginResponse{
				Status: int64(http.StatusUnauthorized),
				Error:  "password-not-match",
			},
		},
		{
			"User Not Found",
			&pb.LoginRequest{
				Username: "xxxx",
				Password: "xxxx",
			},
			&pb.LoginResponse{
				Status: int64(http.StatusNotFound),
				Error:  "user-not-found",
			},
		},
	}

	ctx := context.Background()
	conn := checkConnection(ctx, t)
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			response, err := client.Login(ctx, tc.req)
			require.NoError(t, err)

			require.Equal(t, tc.resp.Status, response.Status)
			require.Equal(t, tc.resp.Error, response.Error)

			if response.Status == int64(http.StatusOK) {
				require.NotEmpty(t, response.Token)
			}
		})
	}
}

// For Company
func TestShowUserRegister(t *testing.T) {
	testCases := []struct {
		name string
		req  *pb.ShowUserRequest
		resp *pb.ShowUserResponse
	}{
		{
			"OK",
			&pb.ShowUserRequest{
				Limit: 10,
				Page:  1,
			},
			&pb.ShowUserResponse{
				Status: int64(http.StatusOK),
				Error:  "",
			},
		},
		{
			"Invalid Limit",
			&pb.ShowUserRequest{
				Limit: 0,
				Page:  1,
			},
			&pb.ShowUserResponse{
				Status: int64(http.StatusBadRequest),
				Error:  "invalid-limit",
			},
		},
		{
			"Invalid Page",
			&pb.ShowUserRequest{
				Limit: 10,
				Page:  0,
			},
			&pb.ShowUserResponse{
				Status: int64(http.StatusBadRequest),
				Error:  "invalid-page",
			},
		},
		{
			"Not Found",
			&pb.ShowUserRequest{
				Limit: 10,
				Page:  99,
			},
			&pb.ShowUserResponse{
				Status: int64(http.StatusNotFound),
				Error:  "user-not-found",
			},
		},
	}

	ctx := context.Background()
	conn := checkConnection(ctx, t)
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			response, err := client.ShowRegisterUser(ctx, tc.req)
			require.NoError(t, err)

			require.Equal(t, tc.resp.Status, response.Status)
			require.Equal(t, tc.resp.Error, response.Error)

			if response.Status == int64(http.StatusOK) {
				require.NotZero(t, len(response.Users))
			}
		})
	}
}

func TestShowAllUserPoint(t *testing.T) {
	testCases := []struct {
		name string
		req  *pb.ShowUserRequest
		resp *pb.ShowUserResponse
	}{
		{
			"OK",
			&pb.ShowUserRequest{
				Limit: 10,
				Page:  1,
			},
			&pb.ShowUserResponse{
				Status: int64(http.StatusOK),
				Error:  "",
			},
		},
		{
			"Invalid Limit",
			&pb.ShowUserRequest{
				Limit: 0,
				Page:  1,
			},
			&pb.ShowUserResponse{
				Status: int64(http.StatusBadRequest),
				Error:  "invalid-limit",
			},
		},
		{
			"Invalid Page",
			&pb.ShowUserRequest{
				Limit: 10,
				Page:  0,
			},
			&pb.ShowUserResponse{
				Status: int64(http.StatusBadRequest),
				Error:  "invalid-page",
			},
		},
		{
			"Not Found",
			&pb.ShowUserRequest{
				Limit: 10,
				Page:  99,
			},
			&pb.ShowUserResponse{
				Status: int64(http.StatusNotFound),
				Error:  "user-not-found",
			},
		},
	}

	ctx := context.Background()
	conn := checkConnection(ctx, t)
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			response, err := client.ShowAllUserPoint(ctx, tc.req)
			require.NoError(t, err)

			require.Equal(t, tc.resp.Status, response.Status)
			require.Equal(t, tc.resp.Error, response.Error)

			if response.Status == int64(http.StatusOK) {
				require.NotZero(t, len(response.Users))
			}
		})
	}
}

func TestSetArticlePoint(t *testing.T) {
	const (
		ARTICLE_ID = 1
		COMPANY_ID = 2
	)
	testCases := []struct {
		name string
		req  *pb.SetPointRequest
		resp *pb.GenericResponse
	}{
		{
			"OK",
			&pb.SetPointRequest{
				ArticleId: ARTICLE_ID,
				UserId:    COMPANY_ID,
				Point:     100,
			},
			&pb.GenericResponse{
				Status: int64(http.StatusOK),
				Error:  "",
			},
		},
		{
			"Invalid Article ID",
			&pb.SetPointRequest{
				ArticleId: 0,
				UserId:    COMPANY_ID,
				Point:     100,
			},
			&pb.GenericResponse{
				Status: int64(http.StatusBadRequest),
				Error:  "invalid-article-id",
			},
		},
		{
			"Invalid Company ID",
			&pb.SetPointRequest{
				ArticleId: ARTICLE_ID,
				UserId:    0,
				Point:     100,
			},
			&pb.GenericResponse{
				Status: int64(http.StatusBadRequest),
				Error:  "invalid-user-id",
			},
		},
		{
			"Invalid Point",
			&pb.SetPointRequest{
				ArticleId: ARTICLE_ID,
				UserId:    COMPANY_ID,
				Point:     0,
			},
			&pb.GenericResponse{
				Status: int64(http.StatusBadRequest),
				Error:  "invalid-point",
			},
		},
		{
			"User Not Found",
			&pb.SetPointRequest{
				ArticleId: ARTICLE_ID,
				UserId:    99999,
				Point:     100,
			},
			&pb.GenericResponse{
				Status: int64(http.StatusNotFound),
				Error:  "user-not-found",
			},
		},
		{
			"Article Not Found",
			&pb.SetPointRequest{
				ArticleId: 99999,
				UserId:    COMPANY_ID,
				Point:     100,
			},
			&pb.GenericResponse{
				Status: int64(http.StatusNotFound),
				Error:  "article-not-found",
			},
		},
	}

	ctx := context.Background()
	conn := checkConnection(ctx, t)
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			response, err := client.SetPoint(ctx, tc.req)
			require.NoError(t, err)

			require.Equal(t, tc.resp.Status, response.Status)
			require.Equal(t, tc.resp.Error, response.Error)
		})
	}

}
