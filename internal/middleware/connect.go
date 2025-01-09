package middleware

import (
	"dragon_battle/db"
	"dragon_battle/internal/untils"
	"log"
)

func Connet_wallet() {
	pl := untils.Crea_new_data()
	err := db.Cre_Player(pl)
	if err != nil {
		log.Fatal("failler connect")
	}
	log.Fatal("connect and create_player")
}
