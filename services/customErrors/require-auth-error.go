package customErrors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAuthErrorResponse(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": ApiError{Message: "Not signed in"}})
	c.Abort()
}
