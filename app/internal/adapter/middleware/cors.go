package middleware

import (
	"net/http"
	"strings"

	"github.com/takusankai/portfolio_tools_hub_backend/app/internal/utils"
)

// CORS はクロスオリジンリクエストを許可するミドルウェア
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 許可するオリジンのリスト（環境変数から取得するか、デフォルト値を使用）
		allowedOriginsStr := utils.GetEnv("ALLOWED_ORIGINS", "http://localhost:3000,https://portfolio-tools-hub-flontend.vercel.app")
		requestOrigin := r.Header.Get("Origin")

		// オリジンが空の場合は何もしない
		if requestOrigin == "" {
			next.ServeHTTP(w, r)
			return
		}

		// カンマで区切られたオリジンを配列に分割
		allowedOrigins := strings.Split(allowedOriginsStr, ",")
		allowed := false

		// オリジンが許可リストに含まれているか確認
		if allowedOriginsStr == "*" {
			allowed = true
		} else {
			for _, allowedOrigin := range allowedOrigins {
				if strings.TrimSpace(allowedOrigin) == requestOrigin {
					allowed = true
					break
				}
			}
		}

		if allowed {
			w.Header().Set("Access-Control-Allow-Origin", requestOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
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
