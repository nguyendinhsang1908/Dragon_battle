package api

import (
	"dragon_battle/internal/game"
	"log"

	"github.com/gin-gonic/gin"
)

func Api() {
	r := gin.Default()

	//Get,Update Profile
	r.GET("/get_Profile", game.GetInfoPlayer_API())
	r.PATCH("update_name", game.UpdateName())
	r.PATCH("/update_ava", game.UpdateAvatar())

	//Add,Get,Delete Friend
	r.POST("/add_friend", game.AddFriend_API())
	r.GET("/get_friend", game.GetFriend_API())
	r.DELETE("/delete_friend", game.DeleteFriendAPI())
	r.POST("/send_fr_request", game.Send_Fr_Mess_Request())
	r.POST("/accept_fr_request", game.Accept_Fr_Request())
	r.POST("/reject_fr_request", game.Reject_Fr_Request())

	//Post,Get Messages
	r.POST("/post_messange_api", game.PostMessage_API())
	r.GET("/get_messange_api", game.GetMessagesAPI())

	//Create clan
	r.POST("/create_clan", game.Crea_Clan())

	//Add,Delete Dragon
	r.POST("/Add_dragon", game.Crate_Dragon())
	r.DELETE("/Delete_dragon", game.Delete_Dragon())
	r.GET("/GetAllDragon", game.GetAllDragon())

	// Chạy server
	if err := r.Run("localhost:8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
