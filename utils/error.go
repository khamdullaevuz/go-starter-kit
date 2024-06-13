package utils

import "go-starter-kit/models"

func NewError(message string, request models.Request) (Error models.Error){
	return models.Error{
		Jsonrpc: request.Jsonrpc,
		Id: request.Id,
		Status: false,
		Message: message,
	}
}