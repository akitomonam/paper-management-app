// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"upload/file" //追加
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file.UploadHandler(w, r) //追加
}

func setupRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload", uploadHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("ファイルアップロード開始")
	setupRoutes()
}
