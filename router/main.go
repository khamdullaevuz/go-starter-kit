package router

import (
	"go-starter-kit/models"
	"go-starter-kit/router/middlewares"
	"go-starter-kit/services"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.Use(middlewares.JsonMiddleware())
		v1.GET("/services", func(c *gin.Context) {
			service := services.NewMethodService()

			response, error := service.SwitchMethods(
				c.MustGet("request").(models.Request),
			)

			if error != nil{
				c.JSON(400, error)
			}else{
				c.JSON(200, response)
			}

		})
	}

	return router
}