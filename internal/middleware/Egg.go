package middleware

import (
	"dragon_battle/db"
	"dragon_battle/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add_egg() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request models.Egg
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gọi hàm DeleteFriend để xóa bạn bè
		err := db.Db_Add_Egg(request.Name, request.Rate)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Add Egg", "details": err.Error()})
			return
		}

		// Trả về thông báo thành công
		g.JSON(http.StatusOK, gin.H{"message": "Add Egg successfully"})
	}
}

func Delete_eggid() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			Egg_id int `json:"egg_id" binding:"required"`
		}
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gọi hàm DeleteFriend để xóa bạn bè
		err := db.Db_Delete_EggID(request.Egg_id)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Delete Egg", "details": err.Error()})
			return
		}

		// Trả về thông báo thành công
		g.JSON(http.StatusOK, gin.H{"message": "Delete Egg successfully"})
	}
}

func Delete_eggName() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			Egg_name string `json:"egg_name" binding:"required"`
		}
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gọi hàm DeleteFriend để xóa bạn bè
		err := db.Db_Delete_EggName(request.Egg_name)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Delete Egg", "details": err.Error()})
			return
		}

		// Trả về thông báo thành công
		g.JSON(http.StatusOK, gin.H{"message": "Delete Egg successfully"})
	}
}

func GetAllEgg() func(g *gin.Context) {
	return func(g *gin.Context) {
		// Gọi hàm
		eggs, err := db.GetAllEggs()
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Get All Egg", "details": err.Error()})
			return
		}

		// Trả về thông báo thành công
		g.JSON(http.StatusOK, gin.H{"message": "Get All Eggs successfully", "Eggs": eggs})
	}
}
