package handler

import (
	"ZeroWebui/config"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func Api(c *gin.Context) {
	action := c.Param("action")

	target := config.Ollama + "/api/" + action

	req, err := http.NewRequest(c.Request.Method, target, c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to create request: %v", err)
		return
	}

	//copy header
	for k, v := range c.Request.Header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}

	//send
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.String(http.StatusBadGateway, "failed to forward: %v", err)
		return
	}
	defer resp.Body.Close()

	//copy response header
	for k, v := range resp.Header {
		for _, vv := range v {
			c.Header(k, vv)
		}
	}

	c.Status(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}
