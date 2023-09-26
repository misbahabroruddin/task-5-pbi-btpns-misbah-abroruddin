// auth_middleware.go

package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/misbahabroruddin/task-5-pbi-btpns-misbah-abroruddin/helpers"
)

// AuthMiddleware is a middleware function that checks the JWT token for authentication.
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			ctx.Abort()
			return
		}

		// Check if the Authorization header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			ctx.Abort()
			return
		}

		// Extract the token from the header
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify the token
		issuer, err := helpers.ParseJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		// You can store the authenticated user's ID or other information in the context if needed
		// if issuer == "" {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		// 	ctx.Abort()
		// 	return
		// }
		ctx.Set("userID", issuer.Id)

		// Continue processing the request
		ctx.Next()
	}
}
