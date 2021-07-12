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
	feature3()
	feature4()

	r.Run()
}

func feature1() {
	fmt.Println("Feature 1 activated")
}

func feature2() {
	fmt.Println("Feature 2 activated")
}

func feature3() {
	fmt.Println("Feature for v1.2.3 activated")
}

func feature4() {
	fmt.Println("Feature for v1.2.3 activated")
}
