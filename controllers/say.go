package controllers

import (
	"go-starter-kit/models"
	"go-starter-kit/utils"
)

type SayController struct{}

func (controller SayController) Hello(request models.Request) models.Response {
	return utils.NewResponse(map[string]interface{}{
		"message": "Hello",
		"params":  request.Params,
	}, request)
}

func (controller SayController) Bye(request models.Request) models.Response {
	return utils.NewResponse(map[string]interface{}{
		"message": "Good Bye",
		"params":  request.Params,
	}, request)
}
