package main

import (
	"argo-pipeline/routes"
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
	fmt.Println("Feature 1 for 1.1.0")
}

func feature2() {
	fmt.Println("Feature 2 for v1.1.0")
}
