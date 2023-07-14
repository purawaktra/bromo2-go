package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckBasicAuth(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		Warn("CheckBasicAuth", "Basic auth not set")
		c.JSON(http.StatusUnauthorized, "")
		c.Abort()
		return
	}
	if username != AppAuthUsername {
		Warn("CheckBasicAuth", "Username given not valid")
		c.JSON(http.StatusUnauthorized, "")
		c.Abort()
		return
	}
	if password != AppAuthPassword {
		Warn("CheckBasicAuth", "Password given not valid")
		c.JSON(http.StatusUnauthorized, "")
		c.Abort()
		return
	}
	Info("Basic auth successful")
	return
}
