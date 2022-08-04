package services

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/maslow123/buddyku-users/pkg/pb"
	"github.com/maslow123/buddyku-users/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateArticle(t *testing.T) {
	testCases := []struct {
		name string
		req  *pb.CreateArticleRequest
		resp *pb.GenericResponse
	}{
		{
			"OK",
			&pb.CreateArticleRequest{
				Title:     fmt.Sprintf("Title %s", utils.RandomString(50)),
				Content:   fmt.Sprintf("Content %s", utils.RandomString(100)),
				CreatedBy: 1,
			},
			&pb.GenericResponse{
				Status: http.StatusCreated,
				Error:  "",
			},
		},
		{
			"Invalid Title",
			&pb.CreateArticleRequest{
				Title:     "",
				Content:   fmt.Sprintf("Content %s", utils.RandomString(100)),
				CreatedBy: 1,
			},
			&pb.GenericResponse{
				Status: http.StatusBadRequest,
				Error:  "invalid-title",
			},
		},
		{
			"Invalid Content",
			&pb.CreateArticleRequest{
				Title:     fmt.Sprintf("Title %s", utils.RandomString(50)),
				Content:   "",
				CreatedBy: 1,
			},
			&pb.GenericResponse{
				Status: http.StatusBadRequest,
				Error:  "invalid-content",
			},
		},
		{
			"Invalid Created By",
			&pb.CreateArticleRequest{
				Title:     fmt.Sprintf("Title %s", utils.RandomString(50)),
				Content:   fmt.Sprintf("Content %s", utils.RandomString(100)),
				CreatedBy: 0,
			},
			&pb.GenericResponse{
				Status: http.StatusBadRequest,
				Error:  "invalid-created-by",
			},
		},
	}

	ctx := context.Background()
	conn := checkConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArticleServiceClient(conn)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			response, err := client.CreateArticle(ctx, tc.req)
			require.NoError(t, err)

			require.Equal(t, tc.resp.Status, response.Status)
			require.Equal(t, tc.resp.Error, response.Error)
		})
	}
}

func TestDetailArticle(t *testing.T) {
	testCases := []struct {
		name      string
		articleID func(t *testing.T, ctx context.Context, client pb.ArticleServiceClient) int32
		resp      *pb.GenericResponse
	}{
		{
			"OK",
			func(t *testing.T, ctx context.Context, client pb.ArticleServiceClient) int32 {
				return createArticle(t, ctx, client)
			},
			&pb.GenericResponse{
				Status: http.StatusOK,
				Error:  "",
			},
		},
		{
			"Invalid ID",
			func(t *testing.T, ctx context.Context, client pb.ArticleServiceClient) int32 {
				return 0
			},
			&pb.GenericResponse{
				Status: http.StatusBadRequest,
				Error:  "invalid-article-id",
			},
		},
		{
			"Article Not Found",
			func(t *testing.T, ctx context.Context, client pb.ArticleServiceClient) int32 {
				return 999999
			},
			&pb.GenericResponse{
				Status: http.StatusNotFound,
				Error:  "article-not-found",
			},
		},
	}

	ctx := context.Background()
	conn := checkConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArticleServiceClient(conn)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			req := &pb.ArticleDetailRequest{
				Id: tc.articleID(t, ctx, client),
			}
			response, err := client.DetailArticle(ctx, req)
			require.NoError(t, err)

			require.Equal(t, tc.resp.Status, response.Status)
			require.Equal(t, tc.resp.Error, response.Error)
		})
	}
}

func TestGetUserArticlePoint(t *testing.T) {
	testCases := []struct {
		name string
		req  *pb.UserArticlePointRequest
		resp *pb.UserArticlePointResponse
	}{
		{
			"OK",
			&pb.UserArticlePointRequest{
				CreatedBy: 1,
			},
			&pb.UserArticlePointResponse{
				Status:     http.StatusOK,
				Error:      "",
				TotalPoint: 0,
			},
		},
		{
			"Invalid Created By",
			&pb.UserArticlePointRequest{
				CreatedBy: 0,
			},
			&pb.UserArticlePointResponse{
				Status: http.StatusBadRequest,
				Error:  "invalid-created-by",
			},
		},
		{
			"Not Found",
			&pb.UserArticlePointRequest{
				CreatedBy: 9999,
			},
			&pb.UserArticlePointResponse{
				Status: http.StatusNotFound,
				Error:  "article-not-found",
			},
		},
	}

	ctx := context.Background()
	conn := checkConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArticleServiceClient(conn)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			response, err := client.GetUserArticlePoint(ctx, tc.req)
			require.NoError(t, err)

			require.Equal(t, tc.resp.Status, response.Status)
			require.Equal(t, tc.resp.Error, response.Error)
		})
	}
}
func createArticle(t *testing.T, ctx context.Context, client pb.ArticleServiceClient) int32 {
	arg := &pb.CreateArticleRequest{
		Title:     fmt.Sprintf("Title %s", utils.RandomString(50)),
		Content:   fmt.Sprintf("Content %s", utils.RandomString(100)),
		CreatedBy: 1,
	}

	response, err := client.CreateArticle(ctx, arg)
	require.NoError(t, err)
	require.NotZero(t, response.Article.Id)

	lastInsertedId := response.Article.Id
	return lastInsertedId
}
