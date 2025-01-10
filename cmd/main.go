package main

import (
	"dragon_battle/api"
	"dragon_battle/config"
	"dragon_battle/db"
	"log"
)

func main() {
	cfg := config.LoadDBConfig()

	db.InitDB(cfg)

	log.Println("Connect_database successfull!")

	api.Api()

}
