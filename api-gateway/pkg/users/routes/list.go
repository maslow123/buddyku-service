package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maslow123/api-gateway/pkg/users/pb"
	"github.com/maslow123/api-gateway/pkg/utils"
)

func ShowRegisterUser(ctx *gin.Context, c pb.UserServiceClient) {

	limitString := ctx.Query("limit")
	pageString := ctx.Query("page")

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	page, err := strconv.Atoi(pageString)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	res, err := c.ShowRegisterUser(context.Background(), &pb.ShowUserRequest{
		Page:  int32(page),
		Limit: int32(limit),
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	utils.SendProtoMessage(ctx, res, int(res.Status))
}
