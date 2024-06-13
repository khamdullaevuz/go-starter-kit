package models

type Response struct {
	Jsonrpc string      `json:"jsonrpc" binding:"required"`
	Id      int64       `json:"id" binding:"required"`
	Status bool      `json:"status" binding:"required"`
	Message string      `json:"message" binding:"required"`
	Data interface{} `json:"data"`
}