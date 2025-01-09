package config

import "os"

type DBConfig struct {
	User   string
	Passwd string
	Net    string
	Addr   string
	DBName string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		User:   os.Getenv("root"),
		Passwd: os.Getenv("123456"),
		Addr:   os.Getenv("127.0.0.1") + ":" + os.Getenv("3307"),
		DBName: os.Getenv("game_roulette"),
	}
}
