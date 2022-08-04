package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maslow123/api-gateway/pkg/users/pb"
	"github.com/maslow123/api-gateway/pkg/utils"
)

type SetArticlePointRequestBody struct {
	ArticleID int32 `json:"article_id"`
	Point     int32 `json:"point"`
}

func SetArticlePoint(ctx *gin.Context, c pb.UserServiceClient) {
	userID := ctx.Value("user_id").(int32)

	req := SetArticlePointRequestBody{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.SetPoint(context.Background(), &pb.SetPointRequest{
		ArticleId: req.ArticleID,
		UserId:    userID,
		Point:     req.Point,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	utils.SendProtoMessage(ctx, res, int(res.Status))
}
