package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//This will help with automated uptime checks
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
