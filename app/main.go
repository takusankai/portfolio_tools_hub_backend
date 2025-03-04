package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/takusankai/portfolio_tools_hub_backend/app/sample_package"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "環境変数:")
	fmt.Fprintf(w, "DB_HOST: %s\n", os.Getenv("DB_HOST"))
	fmt.Fprintf(w, "POSTGRES_USER: %s\n", os.Getenv("POSTGRES_USER"))

	data := sample_package.Data()
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
