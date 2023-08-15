package main

import (
	"Project2/Gin_first/router"
	"log"
)

func main() {
	router := router.InitRouter()
	err := router.Run()
	if err != nil {
		log.Println("gin run error" + err.Error())
	}
}
