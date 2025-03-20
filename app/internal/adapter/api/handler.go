package api

import (
	"fmt"
	"log"
	"net/http"

	generated "github.com/takusankai/portfolio_tools_hub_backend/app/gen"
	"github.com/takusankai/portfolio_tools_hub_backend/app/internal/utils"
)

// Handler はAPIハンドラー構造体
type Handler struct{}

// NewHandler は新しいハンドラーを作成する
func NewHandler() *Handler {
	return &Handler{}
}

// ErrorHandler はAPI呼び出し時のエラーを処理する
func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("API Error: %v", err)

	// エラータイプに基づいてステータスコードを調整
	switch e := err.(type) {
	case *generated.RequiredParamError:
		utils.WriteJSON(w, http.StatusBadRequest, generated.BadRequest{
			Error: fmt.Sprintf("必須パラメータがありません: %s", e.ParamName),
		})
	case *generated.InvalidParamFormatError:
		utils.WriteJSON(w, http.StatusBadRequest, generated.BadRequest{
			Error: fmt.Sprintf("不正なパラメータ形式: %s", e.ParamName),
		})
	default:
		utils.WriteJSON(w, http.StatusInternalServerError, generated.InternalServerError{
			Error: "内部サーバーエラーが発生しました",
		})
	}
}
