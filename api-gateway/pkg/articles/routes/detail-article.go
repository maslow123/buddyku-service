package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maslow123/api-gateway/pkg/articles/pb"
	"github.com/maslow123/api-gateway/pkg/utils"
)

func DetailArticle(ctx *gin.Context, c pb.ArticleServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.DetailArticle(context.Background(), &pb.ArticleDetailRequest{
		Id: int32(id),
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	if res.Status != http.StatusOK {
		utils.SendProtoMessage(ctx, res, int(res.Status))
		return
	}

	utils.SendProtoMessage(ctx, res, int(res.Status))
}
