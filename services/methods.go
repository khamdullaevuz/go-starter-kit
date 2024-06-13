package services

import (
	"go-starter-kit/models"
	"go-starter-kit/utils"
)

type MethodService struct {}

func NewMethodService() MethodService {
	return MethodService{}
}

func (service MethodService) SwitchMethods (request models.Request) (interface{}, interface{}){
	routes := map[string]func(models.Request) (models.Response, interface{}){
		"say.hello": SayHello,
		"say.bye": SayBye,
	}

	if method, ok := routes[request.Method]; ok {
		return method(request)
	}

	return nil, utils.NewError("Method not found", request)
}

func SayHello(request models.Request) (models.Response, interface{}){
	return utils.NewResponse(map[string]interface{}{
			"message": "Hello",
			"params": request.Params,
		}, request), nil
}

func SayBye(request models.Request) (models.Response, interface{}){
	return utils.NewResponse(map[string]interface{}{
			"message": "Good Bye",
			"params": request.Params,
		}, request), nil
}