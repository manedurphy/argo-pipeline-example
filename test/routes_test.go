package test

import (
	"architecture/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/ping", routes.Ping)

	req, err := http.NewRequest(http.MethodGet, "/ping", nil)

	if err != nil {
		t.Errorf("could not make request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "status should be 200")
}

func TestHealthz(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/healthz", routes.Healthz)

	req, err := http.NewRequest(http.MethodGet, "/healthz", nil)

	if err != nil {
		t.Errorf("could not make request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "status should be 200")
}
