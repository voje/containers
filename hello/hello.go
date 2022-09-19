package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"net/http"
)

func getHelloEnv() string {
	val, ok := os.LookupEnv("HELLO")
	if !ok {
		return "Generic hello server"
	}
	return val
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(getHelloEnv()))
	})
	r.Run()
}
