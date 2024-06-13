package models

type Error struct{
	Jsonrpc string      `json:"jsonrpc" binding:"required"`
	Id      int64       `json:"id" binding:"required"`
	Status bool      `json:"status" binding:"required"`
	Error interface{} `json:"error"`
}