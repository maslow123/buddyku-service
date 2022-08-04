package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maslow123/api-gateway/pkg/articles/pb"
	"github.com/maslow123/api-gateway/pkg/utils"
)

func GetUserArticlePoint(ctx *gin.Context, c pb.ArticleServiceClient) {
	userID := ctx.Value("user_id").(int32)

	res, err := c.GetUserArticlePoint(context.Background(), &pb.UserArticlePointRequest{
		CreatedBy: userID,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	utils.SendProtoMessage(ctx, res, int(res.Status))
}
