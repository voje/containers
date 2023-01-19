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

func getHelloPathEnv() string {
	val, ok := os.LookupEnv("HELLO_PATH")
	if !ok {
		return "/"
	}
	return val
}

func main() {
	r := gin.Default()
	r.GET(getHelloPathEnv(), func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(getHelloEnv()))
	})
	r.Run()
}
