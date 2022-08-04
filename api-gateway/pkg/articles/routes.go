package users

import (
	"github.com/gin-gonic/gin"
	"github.com/maslow123/api-gateway/pkg/articles/routes"
	"github.com/maslow123/api-gateway/pkg/config"
	"github.com/maslow123/api-gateway/pkg/users"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, userSvc *users.ServiceClient) *ServiceClient {

	svc := &ServiceClient{
		Client: InitServiceClient(c),
		Router: r,
	}
	a := users.InitAuthMiddleware(userSvc)

	routes := r.Group("/articles")
	routes.GET("/detail/:id", svc.DetailArticle)
	routes.Use(a.CORSMiddleware)
	routes.Use(a.AuthRequired)
	routes.POST("/create", svc.CreateArticle)
	routes.GET("/point", svc.GetUserArticlePoint)

	return svc
}

func (svc *ServiceClient) CreateArticle(ctx *gin.Context) {
	routes.CreateArticle(ctx, svc.Client)
}

func (svc *ServiceClient) DetailArticle(ctx *gin.Context) {
	routes.DetailArticle(ctx, svc.Client)
}

func (svc *ServiceClient) GetUserArticlePoint(ctx *gin.Context) {
	routes.GetUserArticlePoint(ctx, svc.Client)
}
