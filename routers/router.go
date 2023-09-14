package routers

import (
	"api-ppob/app/controllers"
	"api-ppob/app/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.Engine) {
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	route.POST("/login", controllers.LoginUser)
	route.POST("/registration", controllers.RegisterUser)

	auth := route.Group("/")
	auth.Use(middlewares.CheckUser())
	{
		auth.GET("/profile", controllers.ProfileUser)
	}

}
