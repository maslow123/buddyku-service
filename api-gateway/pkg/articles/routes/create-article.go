package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maslow123/api-gateway/pkg/articles/pb"
	"github.com/maslow123/api-gateway/pkg/utils"
)

type CreateArticleRequestBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreateArticle(ctx *gin.Context, c pb.ArticleServiceClient) {
	req := CreateArticleRequestBody{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	userID := ctx.Value("user_id").(int32)

	res, err := c.CreateArticle(context.Background(), &pb.CreateArticleRequest{
		Title:     req.Title,
		Content:   req.Content,
		CreatedBy: userID,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	utils.SendProtoMessage(ctx, res, int(res.Status))
}
