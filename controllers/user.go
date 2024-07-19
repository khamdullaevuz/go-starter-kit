package controllers

import (
	"go-starter-kit/models"
	"go-starter-kit/utils"
)

type UserController struct{}

func (controller UserController) List(request models.Request) models.Response {
	return utils.NewResponse(map[string]interface{}{
		"message": "User List",
		"params":  request.Params,
	}, request)
}
