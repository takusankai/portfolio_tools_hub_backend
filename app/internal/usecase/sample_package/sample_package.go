package sample_package

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/nedpals/supabase-go"
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
		return []string{fmt.Sprintf("データベース接続エラー: %v", err)}
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return []string{fmt.Sprintf("データベース疎通確認エラー: %v", err)}
	}

	// SELECT name FROM users の結果を配列に格納
	rows, err := db.Query("SELECT name FROM users")
	if err != nil {
		return []string{fmt.Sprintf("クエリエラー: %v", err)}
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return []string{fmt.Sprintf("スキャンエラー: %v", err)}
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

var supabaseClient *supabase.Client

func old_Data2() []string {
	// 環境変数から取得
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	// Supabaseクライアントの初期化
	supabaseClient = supabase.CreateClient(supabaseUrl, supabaseKey)

	// Supabaseからユーザーデータを取得
	var users []map[string]interface{}
	err := supabaseClient.DB.From("users").Select("*").Execute(&users)
	if err != nil {
		return []string{fmt.Sprintf("Supabaseクエリエラー: %v", err)}
	}

	// ユーザー名を配列に格納
	var names []string
	for _, user := range users {
		names = append(names, user["name"].(string))
	}

	return names
}

func Data2() []string {
	// 環境変数から取得
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	fmt.Printf("接続先URL: %s\n", supabaseUrl)

	// 完全なエンドポイントURLを構築
	endpoint := fmt.Sprintf("%s/rest/v1/users?select=*", supabaseUrl)

	// HTTPクライアントを作成 (タイムアウトを設定)
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	// HTTPリクエストを作成
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return []string{fmt.Sprintf("リクエスト作成エラー: %v", err)}
	}

	// 必要なヘッダーを設定
	req.Header.Set("apikey", supabaseKey)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", supabaseKey))
	req.Header.Set("Content-Type", "application/json")

	// リクエストを送信
	fmt.Println("Supabaseにリクエスト送信中...")
	resp, err := client.Do(req)
	if err != nil {
		return []string{fmt.Sprintf("Supabaseリクエストエラー: %v", err)}
	}
	defer resp.Body.Close()

	fmt.Printf("ステータスコード: %d\n", resp.StatusCode)

	// 成功以外のステータスコードをチェック
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return []string{fmt.Sprintf("Supabaseエラー: ステータス=%d, レスポンス=%s",
			resp.StatusCode, string(bodyBytes))}
	}

	// レスポンスボディを読み込み
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []string{fmt.Sprintf("レスポンス読み込みエラー: %v", err)}
	}

	// JSONをパース
	var users []map[string]interface{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		return []string{fmt.Sprintf("JSONパースエラー: %v", err)}
	}

	// ユーザー名を配列に格納
	var names []string
	for _, user := range users {
		if name, ok := user["name"].(string); ok {
			names = append(names, name)
		}
	}

	// データが空の場合のメッセージ
	if len(names) == 0 {
		return []string{"ユーザーデータが見つかりませんでした"}
	}

	return names
}
