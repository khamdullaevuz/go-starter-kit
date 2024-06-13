package services

import (
	"api-gateway/models"
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

	return nil, models.Error{
		HttpStatus: 404,
		Data: "Method not found",
	}
}

func SayHello(params models.Request) (models.Response, models.Error){
	return models.Response{
		Jsonrpc: "2.0",
		Id: params.Id,
		Status: "success",
		Message: "Hello",
		Data: params.Params,
	}, models.Error{}
}

func SayGoodbye(params models.Request) (models.Response, models.Error){
	return models.Response{
		Jsonrpc: "2.0",
		Id: params.Id,
		Status: "success",
		Message: "Good Bye",
		Data: params.Params,
	}, models.Error{}
}