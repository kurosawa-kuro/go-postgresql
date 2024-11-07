package main

import (
	"database/sql"
	"fmt"
	"go-postgresql/config"
	"log"

	_ "github.com/lib/pq"
)

type Micropost struct {
	ID    int
	Title string
}

// データベース接続を確立
func connectDB(cfg *config.Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal("データベース接続エラー:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("接続テストエラー:", err)
	}
	fmt.Printf("環境 %sのデータベースに正常に接続されました\n", cfg.AppEnv)

	return db
}

// マイクロポストを取得
func fetchMicroposts(db *sql.DB) []Micropost {
	rows, err := db.Query("SELECT id, title FROM microposts")
	if err != nil {
		log.Fatal("クエリ実行エラー:", err)
	}
	defer rows.Close()

	var microposts []Micropost
	for rows.Next() {
		var post Micropost
		if err := rows.Scan(&post.ID, &post.Title); err != nil {
			log.Fatal("データ取得エラー:", err)
		}
		microposts = append(microposts, post)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("行の反復中にエラーが発生:", err)
	}

	return microposts
}

func main() {
	cfg := config.LoadConfig()
	db := connectDB(cfg)
	defer db.Close()

	microposts := fetchMicroposts(db)
	fmt.Printf("取得したマイクロポスト: %+v\n", microposts)
}
