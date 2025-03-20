package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	generated "github.com/takusankai/portfolio_tools_hub_backend/app/gen"
	"github.com/takusankai/portfolio_tools_hub_backend/app/internal/adapter/api"
	appMiddleware "github.com/takusankai/portfolio_tools_hub_backend/app/internal/adapter/middleware"
	"github.com/takusankai/portfolio_tools_hub_backend/app/internal/utils"
)

func main() {
	// APIハンドラーの初期化
	apiHandler := api.NewHandler()

	// Chiルーターの初期化
	r := chi.NewRouter()

	// ミドルウェアの設定
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(appMiddleware.CORS)

	// レガシーハンドラー（後で削除予定）
	// r.Get("/legacy", api.HelloHandler)

	r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		spec, _ := generated.GetSwagger()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(spec)
	})

	// OpenAPI生成コードに基づくサーバー実装
	r.Mount("/", generated.HandlerWithOptions(apiHandler, generated.ChiServerOptions{
		ErrorHandlerFunc: api.ErrorHandler,
	}))

	// サーバー起動
	port := utils.GetEnv("PORT", "8080")
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
