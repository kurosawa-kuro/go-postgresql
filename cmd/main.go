package main

import (
	"fmt"
	"go-postgresql/internal/config"
	"go-postgresql/internal/db"
)

func main() {
	cfg := config.LoadConfig()
	database := db.ConnectDB(cfg)
	defer database.Close()

	microposts := db.FetchMicroposts(database)
	fmt.Printf("取得したマイクロポスト: %+v\n", microposts)
}
