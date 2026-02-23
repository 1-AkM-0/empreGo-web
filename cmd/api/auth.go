package main

import (
	"context"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func (app *application) authCallbackHandler(c *gin.Context) {
	provider := c.Param("provider")
	req := c.Request
	res := c.Writer

	ctx := context.WithValue(req.Context(), "provider", provider)
	req = req.WithContext(ctx)

	ghUser, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "falha no login" + err.Error()})
		return
	}

	finalUserID, err := app.Models.UserModel.GetOrCreateGithubUser(ghUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "erro ao tentar procurar usuário"})
	}

	session := sessions.Default(c)
	session.Set("userID", finalUserID)

	if err := session.Save(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "erro ao salvar sessão"})
		return
	}

	c.Redirect(http.StatusFound, "/v1/jobs")

}

func (app *application) authBeginHandler(c *gin.Context) {
	provider := c.Param("provider")
	ctx := context.WithValue(c.Request.Context(), "provider", provider)
	c.Request = c.Request.WithContext(ctx)

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (app *application) logoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()

	session.Options(sessions.Options{MaxAge: -1, Path: "/"})
	session.Save()

	gothic.Logout(c.Writer, c.Request)

	c.JSON(http.StatusOK, gin.H{"message": "você saiu com sucesso"})
}
