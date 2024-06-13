package utils

import "go-starter-kit/models"

func NewError(message string, request models.Request) (Error models.Response){
	return models.Response{
		Jsonrpc: request.Jsonrpc,
		Id: request.Id,
		Status: false,
		Error: message,
		Data: nil,
	}
}