package main

import (
	"dragon_battle/api"
	"dragon_battle/config"
	"dragon_battle/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadDBConfig()

	db.InitDB(cfg)

	log.Println("Connect_database successfull!")
	r := gin.Default()
	r.POST("/add_friend", api.AddFriend_API())
	r.GET("/get_friend", api.GetFriend_API())
	r.DELETE("/delete_friend", api.DeleteFriendAPI())
	// Chạy server
	if err := r.Run("localhost:8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
