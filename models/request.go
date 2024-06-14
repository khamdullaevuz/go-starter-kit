package models

type Request struct {
	JsonRpc string      `json:"jsonrpc" binding:"required"`
	Id      int64       `json:"id" binding:"required"`
	Method  string      `json:"method" binding:"required"`
	Params  interface{} `json:"params"`
}
