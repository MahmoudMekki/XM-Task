package authentication

import (
	"github.com/MahmoudMekki/XM-Task/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) <= 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 0, "error": "Unauthorized"})
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]
		token, err := utils.GetToken(tokenString)
		if token.Valid {
			claims := token.Claims.(*utils.AuthCustomClaims)
			c.Set("user_data", claims)
			c.Next()
		} else {
			log.Err(err).Msg(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 0, "error": "Unauthorized"})
		}
	}
}
