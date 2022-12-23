// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"upload/file"

	"github.com/jinzhu/gorm"
)

// File_dbs ファイルのメタ情報
type File_dbs struct {
	ID       int
	Filename string    `json:"filename"`
	Filepath string    `json:"filepath"`
	Updateat time.Time `json:"updateAt"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file.UploadHandler(w, r)
}

func getTableList(db *gorm.DB) ([]string, error) {
	// テーブル一覧を取得するSQL文
	query := "SHOW TABLES"
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 取得したテーブル一覧を格納するスライス
	tables := make([]string, 0)

	// 取得した行を1行ずつ処理する
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	return tables, nil
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "sdkogaken"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "test_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

func apiTablesHandler(w http.ResponseWriter, r *http.Request) {
	// CORSのアクセス制御を行う
	w.Header().Set("Access-Control-Allow-Origin", "*")    // 任意のドメインからのアクセスを許可する
	w.Header().Set("Access-Control-Allow-Methods", "GET") // GETメソッドのみを許可する

	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功(アップロード済みファイル一覧取得)")
	}
	defer db.Close()

	// SQLをログ出力するように設定する
	// db.LogMode(true)

	// file_dbsからすべてのレコードを取得する
	var fileDbs []File_dbs
	if err := db.Find(&fileDbs).Error; err != nil {
		// エラーを出力する
		fmt.Println("エラー:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 生成したJSONデータをコンソールに出力する
	// jsonBytes, _ := json.Marshal(fileDbs)
	// fmt.Println(string(jsonBytes))

	// file_dbsの中身をJSON形式で返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(fileDbs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func setupRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload", uploadHandler)
	mux.HandleFunc("/api/tables", apiTablesHandler)

	if err := http.ListenAndServe(":12345", mux); err != nil {
		log.Fatal(err)
	}
}

func serveVueApp() {
	fs := http.FileServer(http.Dir("./fronted/dist"))
	fmt.Println("Vue.jsアプリケーションをサーブ完了")
	http.Handle("/", fs)
}

func main() {
	fmt.Println("ファイルアップロード開始")
	setupRoutes()
	serveVueApp()
}
