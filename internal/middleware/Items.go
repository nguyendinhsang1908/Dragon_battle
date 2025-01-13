package middleware

import (
	"dragon_battle/db"
	"dragon_battle/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add_Item() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request models.Item

		// Example usage of request variable
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call a function to add the item (assuming db.AddItem exists)
		err := db.Add_Item(request.Name, *request.IsStackable, request.MaxStack)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item", "details": err.Error()})
			return
		}

		g.JSON(http.StatusOK, gin.H{"message": "Item added successfully"})
	}
}

func GetItem() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			Na_item string `json:"Name_item" binding:"required"`
		}
		// Example usage of request variable
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call a function to add the item (assuming db.AddItem exists)
		item, err := db.Get_Item(request.Na_item)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Get item", "details": err.Error()})
			return
		}

		g.JSON(http.StatusOK, gin.H{"message": "Item Get successfully", "item": item})
	}

}

func Delete_Item() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			Name string `json:"name_item" binding:"required"`
		}

		// Example usage of request variable
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call a function to add the item (assuming db.AddItem exists)
		err := db.Delete_Item(request.Name)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Delete item", "details": err.Error()})
			return
		}

		g.JSON(http.StatusOK, gin.H{"message": "Item Delete successfully"})
	}
}
