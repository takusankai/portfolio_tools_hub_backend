package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	openapi_types "github.com/oapi-codegen/runtime/types"
	gen_api "github.com/takusankai/portfolio_tools_hub_backend/app/gen_api"
	"github.com/takusankai/portfolio_tools_hub_backend/app/sample_package"
)

// APIサーバーの実装
type apiServer struct{}

// ルートエンドポイント実装
func (s *apiServer) GetRoot(w http.ResponseWriter, r *http.Request) {
	message := map[string]string{
		"message": "here is api",
	}
	writeJSON(w, http.StatusOK, message)
}

// ユーザーリスト取得エンドポイント実装
func (s *apiServer) GetUserList(w http.ResponseWriter, r *http.Request) {
	// 環境変数次第で Data() と Data2() のどちらかを呼び出す
	var data []string
	if os.Getenv("USE_LOCAL_DB") == "true" {
		data = sample_package.Data()
	} else {
		data = sample_package.Data2()
	}

	// データをAPIの形式に変換
	users := []gen_api.UsersUser{}
	for i, name := range data {
		// メールアドレスを生成して型変換
		emailStr := fmt.Sprintf("%s@example.com", strings.ToLower(strings.Replace(name, " ", ".", -1)))

		// 追加変更: stringをopenapi_types.Email型に変換
		users = append(users, gen_api.UsersUser{
			Id:       fmt.Sprintf("user-%d", i+1),
			Username: name,
			Email:    openapi_types.Email(emailStr),
		})
	}

	response := map[string][]gen_api.UsersUser{
		"users": users,
	}

	writeJSON(w, http.StatusOK, response)
}

// レガシーハンドラー (後で削除予定)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world! (legacy endpoint)")
	fmt.Fprintln(w, "test push next to next")
	fmt.Fprintln(w, "final test")
	// 環境変数次第で Data() と Data2() のどちらかを呼び出す
	var data []string
	if os.Getenv("USE_LOCAL_DB") == "true" {
		data = sample_package.Data()
	} else {
		data = sample_package.Data2()
	}
	for _, name := range data {
		fmt.Fprintln(w, name)
	}
}

// CORSミドルウェア - クロスオリジンリクエストを許可
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 許可するオリジンのリスト（環境変数から取得するか、デフォルト値を使用）
		allowedOrigins := getEnv("ALLOWED_ORIGINS", "http://localhost:3000")
		origin := r.Header.Get("Origin")

		// リクエスト元のオリジンが許可リストに含まれているか確認
		if origin != "" && (allowedOrigins == "*" || strings.Contains(allowedOrigins, origin)) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// プリフライトリクエスト（OPTIONS）の処理
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// 環境変数を取得（設定がなければデフォルト値を返す）
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}

// JSONレスポンスを書き込むヘルパー関数
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}

func main() {
	// Chiルーターの初期化
	r := chi.NewRouter()

	// ミドルウェアの設定
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)

	// レガシーハンドラー（後で削除予定）
	r.Get("/legacy", helloHandler)

	// OpenAPI生成コードに基づくサーバー実装
	apiHandler := &apiServer{}
	r.Mount("/", gen_api.Handler(apiHandler))

	// サーバー起動
	port := getEnv("PORT", "8080")
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
