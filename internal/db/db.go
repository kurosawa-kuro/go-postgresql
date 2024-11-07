package db

import (
	"database/sql"
	"fmt"
	"log"

	"go-postgresql/internal/config"
	"go-postgresql/internal/models"

	_ "github.com/lib/pq"
)

// データベース接続を確立
func ConnectDB(cfg *config.Config) *sql.DB {
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
func FetchMicroposts(db *sql.DB) []models.Micropost {
	rows, err := db.Query("SELECT id, title FROM microposts")
	if err != nil {
		log.Fatal("クエリ実行エラー:", err)
	}
	defer rows.Close()

	var microposts []models.Micropost
	for rows.Next() {
		var post models.Micropost
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
