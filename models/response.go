package models

type Response struct {
	JsonRpc string      `json:"jsonrpc" binding:"required"`
	Id      int64       `json:"id" binding:"required"`
	Status  bool        `json:"status" binding:"required"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}
