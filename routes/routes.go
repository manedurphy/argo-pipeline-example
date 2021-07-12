package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetPost(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		internalServerError(c)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		internalServerError(c)
		return
	}

	var post Post
	json.Unmarshal(body, &post)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "healthy",
	})
}

func internalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "internal server error",
	})
}
