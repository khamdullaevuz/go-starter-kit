package main

import (
	"go-starter-kit/router"
)

func main(){
	r := router.Init()

	r.Run(":8080")
}