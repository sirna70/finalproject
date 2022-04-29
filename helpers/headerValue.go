package helpers

import "github.com/gin-gonic/gin"

func GetContentType(c *gin.Context) string {
	contentType := c.Request.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/json"
	}
	return contentType
}
