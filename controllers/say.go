package say

import (
	"go-starter-kit/models"
	"go-starter-kit/utils"
)

func Hello(request models.Request) models.Response {
	return utils.NewResponse(map[string]interface{}{
		"message": "Hello",
		"params":  request.Params,
	}, request)
}

func Bye(request models.Request) models.Response {
	return utils.NewResponse(map[string]interface{}{
		"message": "Good Bye",
		"params":  request.Params,
	}, request)
}
