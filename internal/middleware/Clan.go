package middleware

import (
	"dragon_battle/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Crea_Clan() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			PlayerID int    `json:"player_id" bindding:"required"`
			N_clan   string `json:"name_clan" bindding:"required"`
		}
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := db.Db_Crea_Clan(request.PlayerID, request.N_clan)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		g.JSON(http.StatusOK, gin.H{"message": "Create Clan successfull!!"})
	}

}
