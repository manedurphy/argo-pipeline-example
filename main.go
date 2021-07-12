package main

import (
	"architecture/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", routes.Ping)
	r.GET("/post/:id", routes.GetPost)
	r.GET("/healthz", routes.Healthz)

	feature1()
	feature2()

	r.Run()
}

func feature1() {
	fmt.Println("Feature 1 activated")
}

func feature2() {
	fmt.Println("Feature 2 activated")
}
