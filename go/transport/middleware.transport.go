package transport

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/database"
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
