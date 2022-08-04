package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maslow123/api-gateway/pkg/articles/pb"
	"github.com/maslow123/api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ArticleServiceClient
	Router *gin.Engine
}

func InitServiceClient(c *config.Config) pb.ArticleServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ArticleServiceUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewArticleServiceClient(cc)
}
