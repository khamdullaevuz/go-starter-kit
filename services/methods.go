package services

import (
	"go-starter-kit/models"
	"go-starter-kit/utils"
)

type MethodService struct{}

func NewMethodService() MethodService {
	return MethodService{}
}

func (service MethodService) SwitchMethods(request models.Request) models.Response {
	routes := map[string]func(models.Request) models.Response{
		"say.hello": SayHello,
		"say.bye":   SayBye,
	}

	if method, ok := routes[request.Method]; ok {
		return method(request)
	}

	return utils.NewError(
		utils.ErrorContent("Method not found", -32701, nil),
		request,
	)
}

func SayHello(request models.Request) models.Response {
	return utils.NewResponse(map[string]interface{}{
		"message": "Hello",
		"params":  request.Params,
	}, request)
}

func SayBye(request models.Request) models.Response {
	return utils.NewResponse(map[string]interface{}{
		"message": "Good Bye",
		"params":  request.Params,
	}, request)
}
