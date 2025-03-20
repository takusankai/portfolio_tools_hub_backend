package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

// GetEnv は環境変数を取得（設定がなければデフォルト値を返す）
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		log.Printf("Environment variable %s set to: %s", key, value)
		return value
	}
	log.Printf("Environment variable %s not set, using default value: %s", key, defaultValue)
	return defaultValue
}

// WriteJSON はJSONレスポンスを書き込むヘルパー関数
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}

// ParseInt は文字列型をint型に変換するヘルパー関数
func ParseInt(s string, defaultValue int) int {
	if s == "" {
		return defaultValue
	}
	val, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	return val
}
