package db

import (
	"database/sql"
	"dragon_battle/config"
	"log"
)

func InitDB(cfg config.DBConfig) {
	//dsn := fmt.Sprintf("%s:%s@%s(%s)/%s",
	//	cfg.User, cfg.Passwd, cfg.Net, cfg.Addr, cfg.DBName)
	dsn := "root:123456@tcp(127.0.0.1:3307)/dragon_battle?parseTime=true"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Database connection is not alive: %v", err)
	}

	log.Println("Connected to MySQL database!")
}
