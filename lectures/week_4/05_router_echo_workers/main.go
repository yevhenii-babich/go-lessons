package main

import (
	"echo-sample/route"
)

func main() {
	router := route.Init()
	if err := router.Start(":8081"); err != nil {
		router.Logger.Fatalf("can't start: %v", err)
	}
}
