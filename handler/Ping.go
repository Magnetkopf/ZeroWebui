package handler

import (
	"ZeroWebui/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Ping(c *gin.Context) {
	resp, err := http.Get(config.Ollama)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
		})
		log.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{
			"status": "pong",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
		})
	}
}
