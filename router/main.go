package router

import (
	"api-gateway/models"
	"api-gateway/router/middlewares"
	"api-gateway/services"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.JsonMiddleware())

	router.GET("/services", func(c *gin.Context) {
		
		service := services.NewMethodService()

		serviceResponse, methodErr := service.SwitchMethods(c.MustGet("request").(models.Request))

		if methodErr.HttpStatus != 0 {
			c.JSON(methodErr.HttpStatus, gin.H{
				"message": methodErr.Data,
			})
		}else{
			c.JSON(200, serviceResponse)
		}

	})

	return router
}