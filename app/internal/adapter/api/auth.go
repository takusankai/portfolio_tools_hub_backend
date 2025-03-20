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
