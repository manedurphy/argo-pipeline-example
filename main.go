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
	feature5()
	feature6()

	r.Run()
}

func feature1() {
	fmt.Println("Feature 1 activated v1.2.1 activated")
}

func feature2() {
	fmt.Println("Feature 2 activated v1.2.2 activated")
}

func feature3() {
	fmt.Println("Feature 3 for v1.2.3 activated")
}

func feature4() {
	fmt.Println("Feature 4 for v1.2.3 activated")
}

func feature5() {
	fmt.Println("Feature 5 for v1.2.3 activated")
}

func feature6() {
	fmt.Println("Feature 6 for v1.2.3 activated")
}
