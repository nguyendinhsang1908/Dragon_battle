package api

import (
	"dragon_battle/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func Api() {
	r := gin.Default()

	//Get,Update Profile
	r.GET("/get_Profile", middleware.GetInfoPlayer_API())
	r.PATCH("update_name", middleware.UpdateName())
	r.PATCH("/update_ava", middleware.UpdateAvatar())

	//Add,Get,Delete Friend
	r.POST("/add_friend", middleware.AddFriend_API())
	r.GET("/get_friend", middleware.GetFriend_API())
	r.DELETE("/delete_friend", middleware.DeleteFriendAPI())
	r.POST("/send_fr_request", middleware.Send_Fr_Mess_Request())
	r.POST("/accept_fr_request", middleware.Accept_Fr_Request())
	r.POST("/reject_fr_request", middleware.Reject_Fr_Request())

	//Post,Get Messages
	r.POST("/post_messange_api", middleware.PostMessage_API())
	r.GET("/get_messange_api", middleware.GetMessagesAPI())

	//Create clan
	r.POST("/create_clan", middleware.Crea_Clan())

	//Add,Delete Dragon
	r.POST("/Add_dragon", middleware.Crea_Dragon())
	r.DELETE("/Delete_dragon", middleware.Delete_Dragon())
	r.GET("/GetAllDragon", middleware.GetAllDragon())

	//Add,Delete Items
	r.POST("/add_item", middleware.Add_Item())
	r.DELETE("/delete_item", middleware.Delete_Item())
	r.GET("/get_item", middleware.GetItem())

	//Create Inventory
	//add item to inventory's player
	r.POST("/create_Inventory", middleware.Add_Inventory())
	r.DELETE("/delete_it_inventory", middleware.Delete_item_inventory())
	r.GET("/GetAllItemsInventory", middleware.GetAllItems_inven())
	// Chạy server
	if err := r.Run("localhost:8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
