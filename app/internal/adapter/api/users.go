package api

import (
	"net/http"
	"os"

	generated "github.com/takusankai/portfolio_tools_hub_backend/app/gen"
	"github.com/takusankai/portfolio_tools_hub_backend/app/internal/usecase/sample_package"
	"github.com/takusankai/portfolio_tools_hub_backend/app/internal/utils"
)

// GetUserIdList はユーザーIDリスト取得エンドポイントを実装する
func (h *Handler) GetUserIdList(w http.ResponseWriter, r *http.Request, params generated.GetUserIdListParams) {
	// 環境変数次第でデータ取得元を切り替え
	var data []string
	if os.Getenv("USE_LOCAL_DB") == "true" {
		data = sample_package.Data()
	} else {
		data = sample_package.Data2()
	}

	// 取得上限の処理
	limit := len(data)
	if params.Limit != nil && *params.Limit < limit {
		limit = *params.Limit
	}
	if limit > len(data) {
		limit = len(data)
	}

	// ユーザーIDリストを生成
	userIds := make([]int, 0, limit)
	for i := 0; i < limit; i++ {
		userIds = append(userIds, i+1)
	}

	// レスポンスの作成
	total := len(data)
	response := generated.GetUserIdListResponse{
		UserIdList: userIds,
		Total:      &total,
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

// GetUserNameList はユーザー名リスト取得エンドポイントを実装する
func (h *Handler) GetUserNameList(w http.ResponseWriter, r *http.Request, params generated.GetUserNameListParams) {
	// 環境変数次第でデータ取得元を切り替え
	var data []string
	if os.Getenv("USE_LOCAL_DB") == "true" {
		data = sample_package.Data()
	} else {
		data = sample_package.Data2()
	}

	// 取得上限の処理
	limit := len(data)
	if params.Limit != nil && *params.Limit < limit {
		limit = *params.Limit
	}
	if limit > len(data) {
		limit = len(data)
	}

	// データを制限に合わせて切り取り
	userNames := data
	if limit < len(userNames) {
		userNames = userNames[:limit]
	}

	// レスポンスの作成
	total := len(data)
	response := generated.GetUserNameListResponse{
		UserNameList: userNames,
		Total:        &total,
	}

	utils.WriteJSON(w, http.StatusOK, response)
}
