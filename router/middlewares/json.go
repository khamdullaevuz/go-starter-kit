package middlewares

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"api-gateway/models"
)

func JsonMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := models.Request{}

		data, err := c.GetRawData()

		if err != nil {
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("Error: %v", err),
			})
		}

		json.Unmarshal(data, &request)

		if request.Jsonrpc != "2.0" {
			c.JSON(400, gin.H{
				"message": "Invalid JSON-RPC version",
			})
			c.Abort()
			return
		}

		if request.Id == 0 {
			c.JSON(400, gin.H{
				"message": "Invalid request ID",
			})
			c.Abort()
			return
		}

		if request.Method == "" {
			c.JSON(400, gin.H{
				"message": "Invalid method",
			})
			c.Abort()
			return
		}

		c.Set("request", request)
		
		c.Next()
	}
}