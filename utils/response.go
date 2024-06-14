package utils

import "go-starter-kit/models"

func NewResponse(data interface{}, request models.Request) (Error models.Response) {
	return models.Response{
		JsonRpc: "2.0",
		Id:      request.Id,
		Status:  true,
		Data:    data,
		Error:   nil,
	}
}
