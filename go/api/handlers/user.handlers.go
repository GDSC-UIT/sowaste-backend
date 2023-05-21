package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	Handler services.UserServices
}

// func (uh *UserHandlers) GetUsers(c *gin.Context) {
// 	uh.Handler.GetUsers(c)
// }

func (uh *UserHandlers) GetAUser(c *gin.Context) {
	uh.Handler.GetUser(c)
}

func (uh *UserHandlers) CreateAUser(c *gin.Context) {
	uh.Handler.CreateUser(c)
}

func (uh *UserHandlers) UpdateAUser(c *gin.Context) {
	uh.Handler.UpdateUser(c)
}

func (uh *UserHandlers) DeleteAUser(c *gin.Context) {
	uh.Handler.DeleteUser(c)
}