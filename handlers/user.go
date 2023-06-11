package handlers

import "github.com/gin-gonic/gin"

func (h Handler) GetUser(c *gin.Context) {
	email := c.GetString("email")
	c.JSON(200, gin.H{"email": email})
}
