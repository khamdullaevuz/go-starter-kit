package router

import (
	"go-starter-kit/models"
	"go-starter-kit/router/middlewares"
	"go-starter-kit/services"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.JsonMiddleware())

	router.GET("/services", func(c *gin.Context) {
		
		service := services.NewMethodService()

		serviceResponse, methodErr := service.SwitchMethods(c.MustGet("request").(models.Request))

		if serviceResponse == nil{
			c.JSON(400, methodErr)
		}else{
			c.JSON(200, serviceResponse)
		}

	})

	return router
}