package router

import (
	"api-gateway/models"
	"api-gateway/services"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.GET("/services", func(c *gin.Context) {
		jsonData, err := c.GetRawData()
		if err != nil {
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("Error: %v", err),
			})
		}

		request := models.Request{}

		json.Unmarshal(jsonData, &request)
		
		service := services.NewMethodService()

		serviceResponse, methodErr := service.SwitchMethods(request)

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