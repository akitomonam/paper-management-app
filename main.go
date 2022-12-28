// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
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

// IndexHandler インデックスページを表示するハンドラ
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

// UploadHandler アップロードを行うハンドラ
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file.UploadHandler(w, r)
}

// GetTableList テーブル一覧を取得する
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

// APITablesHandler file_dbsの中身をJSON形式で返すハンドラ
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

	// file_dbsからすべてのレコードを取得する
	var fileDbs []File_dbs
	if err := db.Find(&fileDbs).Error; err != nil {
		// エラーを出力する
		fmt.Println("エラー:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// file_dbsの中身をJSON形式で返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(fileDbs); err != nil {
		fmt.Println("エラー:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteFileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")    // 任意のドメインからのアクセスを許可する
	w.Header().Set("Access-Control-Allow-Methods", "GET") // GETメソッドのみを許可する
	fmt.Println("deleteFileHandlerが呼び出されました")

	fileID, err := strconv.Atoi(r.URL.Query().Get("fileId"))
	// 削除対象のファイルを特定し、削除する処理を実行する
	fmt.Println("fileID:", fileID)
	fmt.Println("fileID:", reflect.TypeOf(fileID))
	// DBに接続する
	db, err := sqlConnect()
	if err != nil {
		http.Error(w, "DB接続失敗", http.StatusInternalServerError)
		fmt.Println("DB接続失敗")
		return
	}
	defer db.Close()

	// file_dbsテーブルからIDを指定してレコードを取得する
	var fileDb File_dbs
	if err := db.Where("id = ?", fileID).First(&fileDb).Error; err != nil {
		// http.Error(w, "レコードが見つかりません", http.StatusNotFound)
		fmt.Println("レコードが見つかりません")
		fmt.Println("レコードが見つかりません:", err)
		// レスポンスを返す
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"result": "false"})
		return
	}
	// file_dbsテーブルからレコードを削除する
	fmt.Println("filedb:", &fileDb)
	if err := db.Delete(&fileDb).Error; err != nil {
		http.Error(w, "データの削除に失敗しました", http.StatusInternalServerError)
		fmt.Println("データの削除に失敗しました")
		return
	}
	//DBから取得したファイルパスでファイルの削除
	if err := os.Remove(fileDb.Filepath); err != nil {
		http.Error(w, "ファイルの削除に失敗しました", http.StatusInternalServerError)
		fmt.Println("ファイルの削除に失敗しました")
		return
	}
	// レスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": "true"})
	fmt.Println("finish deletefile")
}

// pdfのプレビュー処理
func apiPreviewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("apiPreviewHandlerが呼び出されました")
	// CORSのアクセス制御を行う
	w.Header().Set("Access-Control-Allow-Origin", "*")    // 任意のドメインからのアクセスを許可する
	w.Header().Set("Access-Control-Allow-Methods", "GET") // GETメソッドのみを許可する

	// ファイル名を取得する
	fileName := r.URL.Query().Get("fileName")
	fmt.Println("クリックしたファイルの名前:", fileName)

	// ファイルのURLを取得する
	fileUrl, err := getFileUrl(fileName)
	fmt.Println("fileUrl:", fileUrl)
	if err != nil {
		// エラーを出力する
		fmt.Println("エラー:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// レスポンスボディを作成する
		responseBody, err := json.Marshal(map[string]string{
			"fileUrl": fileUrl,
		})
		if err != nil {
			// エラーを出力する
			fmt.Println("エラー:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// レスポンスを送信する
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBody)
	}
}

// ファイルのURLを取得する関数
func getFileUrl(fileName string) (string, error) {
	// ファイルのURLを格納する変数
	var fileUrl string

	// データベースに接続する
	db, err := sqlConnect()
	if err != nil {
		return "", err
	}
	defer db.Close()

	// file_dbsから指定したファイル名のレコードを1件取得する
	var fileDb File_dbs
	// if err := db.Where("filename = ?", fileName).First(&fileDb).Error; err != nil {
	if err := db.Where("filepath = ?", fileName).First(&fileDb).Error; err != nil {
		fmt.Println("エラー(getFileUrl):", err)
		return "", err
	}

	// 取得したレコードのfilepathをfileUrlに代入する
	fileUrl = fileDb.Filepath

	return fileUrl, nil
}

func setupRoutes() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload/file", uploadHandler)
	mux.HandleFunc("/api/delete", deleteFileHandler)
	mux.HandleFunc("/api/tables", apiTablesHandler)
	mux.HandleFunc("/api/preview", apiPreviewHandler)
	mux.Handle("/uploadfiles/", http.StripPrefix("/uploadfiles/", http.FileServer(http.Dir("./uploadfiles"))))

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
