package customErrors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

func MsgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return "Field requires 8 characters"
	case "notFound":
		return "User not found"
	}

	return "error"
}

func GetErrorMessages(err error) []ApiError {
	ve := err.(validator.ValidationErrors)
	out := make([]ApiError, 0)
	for _, fe := range ve {
		out = append(out, ApiError{Message: MsgForTag(fe), Field: fe.Field()})
	}

	return out
}

func NoAuthAllowedResponse(c *gin.Context) {
	errs := make([]ApiError, 0)
	errs = append(errs, ApiError{Message: "User already signed in"})

	c.JSON(http.StatusForbidden, gin.H{"errors": errs})
}
