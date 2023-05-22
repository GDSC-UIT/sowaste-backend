package transport

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/GDSC-UIT/sowaste-backend/go/internal/database"
	"github.com/gin-gonic/gin"

	"firebase.google.com/go/auth"
)

func Recover(sc database.Instance) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")
				return
			}
		}()
		c.Next()
	}
}

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	c.Header("Access-Control-Expose-Headers", "Content-Length")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Max-Age", "86400")
	c.Header("Access-Control-Request-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	c.Header("Access-Control-Request-Method", "GET, POST, PUT, DELETE, OPTIONS")
	c.Next()
}

// AuthMiddleware : to verify all authorized operations
func AuthMiddleware(c *gin.Context) {
	firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)
	var temp *auth.FirebaseInfo
	print(temp.Identities)
	authorizationToken := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	if idToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id token not available"})
		c.Abort()
		return
	}
	//verify token
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	fmt.Println(err.Error())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}
	c.Set("UUID", token.UID)
	c.Next()
}
