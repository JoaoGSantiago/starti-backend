package handlers

import "github.com/gin-gonic/gin"

func errorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func successResponse(c *gin.Context, status int, data any) {
	c.JSON(status, data)
}
