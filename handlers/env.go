package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func (h Handler) Env(c *gin.Context) {
	c.JSON(200, gin.H{
		"env": os.Getenv("ENV"),
	})
}
