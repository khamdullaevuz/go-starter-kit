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
			c.JSON(400, utils.NewError(
				utils.ErrorContent(
					fmt.Sprintf("Invalid request: %s", err.Error()),
					-32700,
					nil,
				), request),
			)
			c.Abort()
			return
		}

		_ = json.Unmarshal(data, &request)

		if request.JsonRpc != "2.0" {
			c.JSON(400, utils.NewError(
				utils.ErrorContent("Invalid JSON-RPC version", -32002, nil),
				request),
			)
			c.Abort()
			return
		}

		if request.Id == 0 {
			c.JSON(422, utils.NewError(
				utils.ErrorContent("Invalid request id", -32003, nil),
				request),
			)
			c.Abort()
			return
		}

		if request.Method == "" {

			c.JSON(404, utils.NewError(
				utils.ErrorContent("Method Not Found", -32004, nil),
				request),
			)
			c.Abort()
			return
		}

		c.Set("request", request)

		c.Next()
	}
}
