package transport

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/GDSC-UIT/sowaste-backend/go/internal/database"
	"github.com/GDSC-UIT/sowaste-backend/go/utils"
	"github.com/gin-gonic/gin"
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
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

// AuthMiddleware : to verify all authorized operations

func AuthMiddleware(c *gin.Context) {
	authorizationToken := c.GetHeader("Authorization")
	fmt.Println(c.Request.Header)
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	fmt.Println(idToken)
	if idToken == "" {
		c.JSON(http.StatusUnauthorized, utils.FailedResponse("Unauthorized"))
		c.Abort()
		return
	}
	claims, isValidate := utils.CheckJwtIsValidateAndGetJwtDecoded(idToken)
	if !isValidate {
		c.JSON(http.StatusUnauthorized, utils.FailedResponse("Unauthorized"))
		c.Abort()
		return
	} else {
		c.Set("CLAIMS", claims)
	}
	c.Next()
}
