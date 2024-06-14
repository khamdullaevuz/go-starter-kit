package middlewares

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"go-starter-kit/models"
	"go-starter-kit/utils"
)

func JsonMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := models.Request{}

		data, err := c.GetRawData()

		if err != nil {
			c.JSON(400, utils.NewError(fmt.Sprintf("Invalid request: %s", err.Error()), request))
			c.Abort()
			return
		}

		json.Unmarshal(data, &request)

		if request.JsonRpc != "2.0" {
			c.JSON(400, utils.NewError("Invalid JSON-RPC version", request))
			c.Abort()
			return
		}

		if request.Id == 0 {
			c.JSON(422, utils.NewError("Invalid request id", request))
			c.Abort()
			return
		}

		if request.Method == "" {

			c.JSON(404, utils.NewError("Method Not Found", request))
			c.Abort()
			return
		}

		c.Set("request", request)

		c.Next()
	}
}
