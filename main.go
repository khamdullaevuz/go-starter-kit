package main

import (
	"api-gateway/router"
)

func main(){
	r := router.Init()

	r.Run(":8080")
}