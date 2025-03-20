package api

import (
	"net/http"
	"time"

	generated "github.com/takusankai/portfolio_tools_hub_backend/app/gen"
	"github.com/takusankai/portfolio_tools_hub_backend/app/internal/utils"
)

// CheckRoot はルートエンドポイントを実装する
func (h *Handler) CheckRoot(w http.ResponseWriter, r *http.Request, params generated.CheckRootParams) {
	// パラメータの処理
	queryValue := "---入力はありません---"
	if params.SampleQuery != nil {
		queryValue = *params.SampleQuery
	}

	// レスポンスの作成
	response := generated.CheckRootResponse{
		EchoQuery:     queryValue,
		SimpleMessage: "APIサーバーが正常に動作しています",
		Timestamp:     time.Now(),
	}

	utils.WriteJSON(w, http.StatusOK, response)
}
