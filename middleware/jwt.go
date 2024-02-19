package middleware

import (
	"gin-socmed/lib"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			lib.HandleError(c, &lib.UnauthorizedError{Message: "Unauthorize"})
			c.Abort()
			return
		}

		userID, err := lib.ValidateToken(tokenString)

		if err != nil {
			lib.HandleError(c, &lib.UnauthorizedError{Message: err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}

}
