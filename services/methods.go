package services

import (
	say "go-starter-kit/controllers"
	"go-starter-kit/models"
	"go-starter-kit/utils"
)

type MethodService struct{}

func NewMethodService() MethodService {
	return MethodService{}
}

func (service MethodService) SwitchMethods(request models.Request) models.Response {
	routes := map[string]func(models.Request) models.Response{
		"say.hello": say.Hello,
		"say.bye":   say.Bye,
	}

	if method, ok := routes[request.Method]; ok {
		return method(request)
	}

	return utils.NewError(
		utils.ErrorContent("Method not found", -32701, nil),
		request,
	)
}
