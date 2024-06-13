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
	switch request.Method {
		case "SayHello":
			return SayHello(request)
		case "SayBye":
			return SayBye(request)
	}

	return nil, utils.NewError("Method not found", request)
}

func SayHello(params models.Request) (models.Response, interface{}){
	return utils.NewResponse(map[string]interface{}{
			"message": "Hello",
			"params": params.Params,
		}, params), nil
}

func SayBye(params models.Request) (models.Response, interface{}){
	return utils.NewResponse(map[string]interface{}{
			"message": "Good Bye",
			"params": params.Params,
		}, params), nil
}