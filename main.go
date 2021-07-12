package main

import (
	"architecture/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", routes.Ping)
	r.GET("/healthz", routes.Healthz)

	r.Run()
}
