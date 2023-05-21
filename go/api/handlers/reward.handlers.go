package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type RewardHandlers struct {
	Handler services.RewardServices
}

func (rh *RewardHandlers) GetRewards(c *gin.Context) {
	rh.Handler.GetRewards(c)
}

func (rh *RewardHandlers) GetAReward(c *gin.Context) {
	rh.Handler.GetAReward(c)
}

func (rh *RewardHandlers) CreateAReward(c *gin.Context) {
	rh.Handler.CreateReward(c)
}

func (rh *RewardHandlers) UpdateAReward(c *gin.Context) {
	rh.Handler.UpdateReward(c)
}

func (rh *RewardHandlers) DeleteAReward(c *gin.Context) {
	rh.Handler.DeleteReward(c)
}
