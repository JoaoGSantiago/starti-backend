package middleware

import (
	"net/http"
	"strings"

	"github.com/JoaoGSantiago/starti-backend/internal/services"
	"github.com/gin-gonic/gin"
)

const UserIDKey = "userID"

func Auth(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro": "autorizaçao invalida"})
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")
		claims, err := jwtService.Validate(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro": "token expirado ou invalido"})
			return
		}

		// Injeta o userID no contexto para os handlers acessarem via c.Get(UserIDKey)
		c.Set(UserIDKey, claims.UserID)
		c.Next()
	}
}
