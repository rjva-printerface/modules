package cookies

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rjva-printerface/auth-service-go/helpers"
)

type CookieSession struct {
	log *helpers.Log
}

func NewCookieSession(log *helpers.Log) *CookieSession {
	return &CookieSession{log}
}

func (c *CookieSession) Create(d interface{}, ctx *gin.Context) {
	ret, err := json.Marshal(d)

	if err != nil {
		c.log.Print(err.Error(), helpers.Red)
	}

	b64 := base64.StdEncoding.EncodeToString(ret)

	ctx.SetCookie("session", b64, 60*60*24, "/", "ticketing-go.dev", true, true)
}

func (c *CookieSession) Get(ctx *gin.Context) []byte {
	cookie, err := ctx.Request.Cookie("session")

	if err != nil {
		return make([]byte, 0)
	}

	jsond, _ := base64.StdEncoding.DecodeString(string(cookie.Value))

	noQuotes := strings.ReplaceAll(string(jsond), string('"'), "")
	value := strings.ReplaceAll(noQuotes, "\"", "")

	return []byte(value)
}

func (c *CookieSession) Distroy(ctx *gin.Context) {
	ctx.SetCookie("session", "", -1, "/", "ticketing-go.dev", true, true)
}

func (c *CookieSession) HasSession(ctx *gin.Context) bool {
	cookie, err := ctx.Request.Cookie("session")

	if err != nil {
		return false
	}

	base64, err := base64.StdEncoding.DecodeString(string(cookie.Value))
	_ = base64

	if err != nil {
		c.log.Print("invalid session", helpers.Red)
		return false
	}

	return true
}
