package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rjva-printerface/auth-service-go/models"
	"github.com/rjva-printerface/auth-service-go/repositories"
	"github.com/rjva-printerface/auth-service-go/services/cookies"
	"github.com/rjva-printerface/auth-service-go/services/customErrors"
	"github.com/rjva-printerface/auth-service-go/services/jwttoken"
)

func RequireAuth(
	session *cookies.CookieSession,
	jwt *jwttoken.JwtToken,
	userRepository *repositories.UserRepository,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !session.HasSession(c) {
			customErrors.RequireAuthErrorResponse(c)
			return
		}

		token := session.Get(c)

		var usrModel models.UserModel
		usr := jwt.DecodeToken(token, &usrModel)

		if usr.Email == "" {
			customErrors.RequireAuthErrorResponse(c)
			return
		}

		findUsr := userRepository.FindById(usr.ID.Hex())

		if findUsr == nil {
			customErrors.RequireAuthErrorResponse(c)
			return
		}

		c.Next()
	}
}
