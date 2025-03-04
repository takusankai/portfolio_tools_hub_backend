package sample_package

import (
	"database/sql"
	"fmt"
	"os"
)

func Data() []string {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("POSTGRES_USER", "portfolio_user")
	password := getEnv("POSTGRES_PASSWORD", "portfolio_password")
	dbname := getEnv("POSTGRES_DB", "portfolio_db")

	// 接続文字列
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// データベース接続
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Sprintf("データベース接続エラー: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return fmt.Sprintf("データベース疎通確認エラー: %v", err)
	}

	// SELECT name FROM users の結果を配列に格納
	rows, err := db.Query("SELECT name FROM users")
	if err != nil {
		return fmt.Sprintf("クエリエラー: %v", err)
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return fmt.Sprintf("スキャンエラー: %v", err)
		}
		names = append(names, name)
	}

	return names
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
