package services

import (
	"go-starter-kit/models"
	"go-starter-kit/utils"
)

type MethodService struct {}

func NewMethodService() MethodService {
	return MethodService{}
}

func (service MethodService) SwitchMethods (request models.Request) (interface{}, models.Error){
	switch request.Method {
		case "api.hello":
			return SayHello(request)
		case "api.goodbye":
			return SayGoodbye(request)
	}

	return nil, utils.NewError("Method not found", request)
}

func SayHello(params models.Request) (models.Response, models.Error){
	return models.Response{
		Jsonrpc: "2.0",
		Id: params.Id,
		Status: true,
		Data: map[string]interface{}{
			"message": "Hello",
			"params": params.Params,
		},
	}, models.Error{}
}

func SayGoodbye(params models.Request) (models.Response, models.Error){
	return models.Response{
		Jsonrpc: "2.0",
		Id: params.Id,
		Status: true,
		Data: map[string]interface{}{
			"message": "Good Bye",
			"params": params.Params,
		},
	}, models.Error{}
}