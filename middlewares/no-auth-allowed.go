package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rjva-printerface/auth-service-go/repositories"
	"github.com/rjva-printerface/auth-service-go/services/cookies"
	"github.com/rjva-printerface/auth-service-go/services/customErrors"
	"github.com/rjva-printerface/auth-service-go/services/jwttoken"
)

func NoAuthAllowed(
	session *cookies.CookieSession,
	jwt *jwttoken.JwtToken,
	userRepository *repositories.UserRepository,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		if session.HasSession(c) {
			customErrors.NoAuthAllowedResponse(c)
			c.Abort()
			return
		}
	}
}
