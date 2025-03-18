package responses

import (
	"path"
	"strings"
)

// PathToRawSpec は外部参照解決のためのヘルパー関数
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))

	// responses.yamlの内容をハードコード...
	responsesYAML := []byte(`...エラーレスポンスの内容...`)

	normalizedPath := path.Clean(pathToFile)
	if strings.HasSuffix(normalizedPath, "responses.yaml") {
		res[normalizedPath] = func() ([]byte, error) {
			return responsesYAML, nil
		}
	}

	return res
}
