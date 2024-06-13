package utils

import "go-starter-kit/models"

func NewError(message string) (Error models.Error){
	return models.Error{
		Jsonrpc: "2.0",
		Id: 0,
		Status: false,
		Message: message,
	}
}