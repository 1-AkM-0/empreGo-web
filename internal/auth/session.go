package auth

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	gorillaSession "github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

func Setup(githubClientID, githubClientSecret, sessionSecret string) cookie.Store {

	gothic.Store = gorillaSession.NewCookieStore([]byte(sessionSecret))

	store := cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400 * 7,
	})

	goth.UseProviders(
		github.New(githubClientID, githubClientSecret, os.Getenv("DOMAIN")+"/v1/auth/github/callback", "user:email"),
	)

	return store
}
