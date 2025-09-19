package handler

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func Index(c *gin.Context) {
	indexPath := filepath.Join("./public", "index.html")
	c.File(indexPath)
}
