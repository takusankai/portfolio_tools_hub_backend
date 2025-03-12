package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/takusankai/portfolio_tools_hub_backend/app/sample_package"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
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

func main() {
	http.HandleFunc("/", helloHandler)

	port := ":8080"
	fmt.Println("Server starting on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
