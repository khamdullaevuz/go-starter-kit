package utils

import "go-starter-kit/models"

func NewResponse(data interface{}, request models.Request) (Error models.Response){
	return models.Response{
		Jsonrpc: request.Jsonrpc,
		Id: request.Id,
		Status: true,
		Data: data,
		Error: nil,
	}
}