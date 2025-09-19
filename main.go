package main

import (
	"ZeroWebui/handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const OllamaAPI = "http://127.0.0.1:11434"

func main() {

	r := gin.Default()
	r.StaticFS("/static", http.Dir("./public/static"))

	r.GET("/api/ping", func(c *gin.Context) {
		handler.Ping(c)
	})

	r.Any("/api/:action", func(c *gin.Context) {
		handler.Api(c)
	})

	r.GET("/", func(c *gin.Context) {
		handler.Index(c)
	})

	if err := r.Run(":9090"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
