package middleware

import (
	"dragon_battle/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateName() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			PlayerID int    `json:"player_id" bindind:"required"`
			Name     string `json:"name" bindding:"required"`
		}
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := db.UpdateName(request.PlayerID, request.Name)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		g.JSON(http.StatusOK, gin.H{"message": "Update name successfull!!"})
	}
}

func UpdateAvatar() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			PlayerID int    `json:"player_id" bindind:"required"`
			Avatar   string `json:"avatar" bindding:"required"`
		}
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := db.UpdateAvatar(request.PlayerID, request.Avatar)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		g.JSON(http.StatusOK, gin.H{"message": "Update avatar successfull!!"})
	}
}

func GetInfoPlayer_API() func(c *gin.Context) {
	return func(c *gin.Context) {
		var request struct {
			PlayerID int `json:"player_id" binding:"required"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		pl, err := db.GetPOnePlayer(request.PlayerID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "Player Information retrieved successfully", "Information": pl})

	}

}
