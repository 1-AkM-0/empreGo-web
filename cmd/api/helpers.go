package main

import (
	"net/url"
	"strconv"

	"github.com/1-AkM-0/empreGo-web/internal/validator"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type envelope map[string]any

func (app *application) readInt(qs url.Values, key string, defaultValue int, v *validator.Validator) int {
	value := qs.Get(key)

	if value == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		v.AddError(key, "deve ser um valor inteiro")
		return defaultValue
	}

	return i
}

func (app *application) tryExtractUserID(c *gin.Context) string {
	session := sessions.Default(c)

	userID, ok := session.Get("userID").(string)
	if ok {
		return userID
	}
	return ""
}
