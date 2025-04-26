package api

import (
	"net/http"
	"time"

	generated "github.com/takusankai/portfolio_tools_hub_backend/app/gen"
	"github.com/takusankai/portfolio_tools_hub_backend/app/internal/utils"
)

// SignUp はサインアップエンドポイントを実装する
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request, params generated.SignUpParams) {
	// パラメータの処理
	queryValue := "クエリなし"
	if params.SampleQuery != nil {
		queryValue = *params.SampleQuery
	}

	// レスポンスの作成
	response := generated.SignUpResponse{
		EchoQuery:     queryValue,
		SimpleMessage: "サインアップリクエストを受け付けました",
		Timestamp:     time.Now(),
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

// Login はユーザーログイン処理を行うハンドラー
func (h *Handler) Login(w http.ResponseWriter, r *http.Request, params generated.LoginParams) {
	// 仮実装（開発中）
	resp := generated.LoginResponse{
		SimpleMessage: "ログイン機能は開発中です",
		EchoQuery:     "",
		Timestamp:     time.Now(),
	}

	// クエリパラメータがあれば設定
	if params.SampleQuery != nil {
		resp.EchoQuery = *params.SampleQuery
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}

// Logout はユーザーログアウト処理を行うハンドラー
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request, params generated.LogoutParams) {
	// 仮実装（開発中）
	resp := generated.LogoutResponse{
		SimpleMessage: "ログアウト機能は開発中です",
		EchoQuery:     "",
		Timestamp:     time.Now(),
	}

	// クエリパラメータがあれば設定
	if params.SampleQuery != nil {
		resp.EchoQuery = *params.SampleQuery
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}
