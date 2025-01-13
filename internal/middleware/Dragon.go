package middleware

import (
	"dragon_battle/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Crea_Dragon() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			Name    string `json:"name_dragon" binding:"required"`
			Point   int    `json:"point" binding:"required"`
			Hp      int    `json:"hp" binding:"required"`
			Mp      int    `json:"mp" binding:"required"`
			Dame    int    `json:"dame" binding:"required"`
			Defense int    `json:"defense" binding:"required"`
			Speed   int    `json:"speed" binding:"required"`
			Rarity  string `json:"rarity" binding:"required"`  // Common,epic,legendary,immo
			Element string `json:"element" binding:"required"` // Fire,Natural,Terra,Ice,Metal,Dark,Light,Chaos
		}

		// Kiểm tra dữ liệu đầu vào
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gọi hàm DeleteFriend để xóa bạn bè
		err := db.Crea_Dragon(request.Name, request.Point, request.Hp, request.Mp, request.Dame, request.Defense, request.Speed, request.Rarity, request.Element)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Add Dragon", "details": err.Error()})
			return
		}

		// Trả về thông báo thành công
		g.JSON(http.StatusOK, gin.H{"message": "Add Dragon successfully"})
	}
}

func Delete_Dragon() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			Name string `json:"name_dr" binding:"required"`
		}

		// Kiểm tra dữ liệu đầu vào
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gọi hàm DeleteFriend để xóa bạn bè
		err := db.Delete_Dragon(request.Name)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Add Dragon", "details": err.Error()})
			return
		}

		// Trả về thông báo thành công
		g.JSON(http.StatusOK, gin.H{"message": "Add Dragon successfully"})
	}
}

func GetAllDragon() func(g *gin.Context) {
	return func(g *gin.Context) {
		// Gọi hàm DeleteFriend để xóa bạn bè
		dragons, err := db.GetAllDragon()
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Get Add Dragon", "details": err.Error()})
			return
		}

		// Trả về thông báo thành công
		g.JSON(http.StatusOK, gin.H{"message": "Get All Dragons successfully", "dragons": dragons})
	}
}
