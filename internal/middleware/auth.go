package middleware

import (
	"net/http"

	"github.com/1-AkM-0/empreGo-web/internal/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RequireAuth(userModel models.UserModel) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")
		if userID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "autenticação necessária"})
			return
		}

		_, err := userModel.GetByID(userID.(string))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "autenticação necessária"})
			return
		}

		c.Next()

	}
}
