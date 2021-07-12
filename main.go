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

	feature2()

	r.Run()
}

func feature2() {
	fmt.Println("Feature 2 for v1.1.0")
}
