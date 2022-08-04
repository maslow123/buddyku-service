package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/maslow123/api-gateway/docs"
	articles "github.com/maslow123/api-gateway/pkg/articles"
	"github.com/maslow123/api-gateway/pkg/config"
	"github.com/maslow123/api-gateway/pkg/users"
)

func main() {
	c, err := config.LoadConfig("./pkg/config/envs", "dev")

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	// Route => handler user
	userService := *users.RegisterRoutes(r, &c)

	// Route => handler articles
	_ = *articles.RegisterRoutes(r, &c, &userService)

	// Start Server
	r.Run(c.Port)
}
