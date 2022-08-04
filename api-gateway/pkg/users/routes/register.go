package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maslow123/api-gateway/pkg/users/pb"
	"github.com/maslow123/api-gateway/pkg/utils"
)

// RegisterRequestBody is represents body of Register request.
type RegisterRequestBody struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Register handles incoming Register Service requests
func Register(ctx *gin.Context, c pb.UserServiceClient) {
	req := RegisterRequestBody{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		ctx.JSON(int(res.Status), utils.ErrorResponse(err))
		return
	}

	ctx.JSON(int(res.Status), &res)
}
