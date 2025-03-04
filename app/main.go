package main

import (
	"fmt"
	"net/http"

	"github.com/takusankai/portfolio_tools_hub_backend/app/sample_package"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	data := sample_package.Data()
	for _, name := range data {
		fmt.Fprintln(w, "Hello,", name)
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
