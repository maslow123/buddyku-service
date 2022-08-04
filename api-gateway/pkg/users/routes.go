package users

import (
	"github.com/gin-gonic/gin"
	"github.com/maslow123/api-gateway/pkg/config"
	"github.com/maslow123/api-gateway/pkg/users/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {

	svc := &ServiceClient{
		Client: InitServiceClient(c),
		Router: r,
	}
	a := InitAuthMiddleware(svc)

	routes := r.Group("/users")
	routes.Use(a.CORSMiddleware)
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	routes.Use(a.AuthRequired)
	routes.GET("/list", svc.ShowRegisterUser)
	routes.GET("/point", svc.ShowAllUserPoint)
	routes.PUT("/set-point", svc.SetArticlePoint)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

// For Company
func (svc *ServiceClient) ShowRegisterUser(ctx *gin.Context) {
	routes.ShowRegisterUser(ctx, svc.Client)
}

func (svc *ServiceClient) ShowAllUserPoint(ctx *gin.Context) {
	routes.ShowAllUserPoint(ctx, svc.Client)
}

func (svc *ServiceClient) SetArticlePoint(ctx *gin.Context) {
	routes.SetArticlePoint(ctx, svc.Client)
}
