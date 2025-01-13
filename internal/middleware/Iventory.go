package middleware

import (
	"dragon_battle/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add_Inventory() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			PlayerID int `json:"player_id" binding:"required"`
			Item     int `json:"item_id" binding:"required"`
			Quantity int `json:"quantity" binding:"required"`
		}
		// Example usage of request variable
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := db.Db_Add_Inventory(request.PlayerID, request.Item, request.Quantity)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Create Inventory", "details": err.Error()})
			return
		}

		g.JSON(http.StatusOK, gin.H{"message": "Create Inventory successfully"})

	}
}

func Delete_item_inventory() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			Inven_id int `json:"inventory_id" binding:"required"`
		}
		// Example usage of request variable
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := db.Db_Delete_item_inventory(request.Inven_id)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Delete Item Inventory", "details": err.Error()})
			return
		}

		g.JSON(http.StatusOK, gin.H{"message": "Delete Inventory item successfully"})

	}
}

func GetAllItems_inven() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			Player_id int `json:"player_id" binding:"required"`
		}
		// Example usage of request variable
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		items, err := db.Db_GetAllItems_inven(request.Player_id)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Get All Item Inventory", "details": err.Error()})
			return
		}

		g.JSON(http.StatusOK, gin.H{"message": "Get All Inventory item successfully", "items": items})

	}
}

func UpdateItems_inven() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			Inven_id int `json:"Inven_id" binding:"required"`
			Item_id  int `json:"item_id" binding:"required"`
			Quantity int `json:"quantity" binding:"required"`
		}
		// Example usage of request variable
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := db.Db_Update_item_inventory(request.Inven_id, request.Item_id, request.Quantity)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Update Item Inventory", "details": err.Error()})
			return
		}

		g.JSON(http.StatusOK, gin.H{"message": "Update Inventory item successfully"})

	}
}
