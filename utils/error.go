package utils

import "go-starter-kit/models"

func NewError(error models.Error, request models.Request) (Error models.Response) {
	return models.Response{
		JsonRpc: request.JsonRpc,
		Id:      request.Id,
		Status:  false,
		Error:   error,
		Data:    nil,
	}
}

func ErrorContent(message string, code int, data interface{}) models.Error {
	return models.Error{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
