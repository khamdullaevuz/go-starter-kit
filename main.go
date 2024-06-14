package main

import (
	"go-starter-kit/router"
)

func main() {
	r := router.Init()

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
