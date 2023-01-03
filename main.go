// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

type Config struct {
	DBMS   string `json:"dbms"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Server string `json:"server"`
	DBName string `json:"dbname"`
	GoPort string `json:"go_port"`
}

type UsersLoginfo struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

var db *gorm.DB

func getConfig() Config {
	// 設定ファイルを開く
	file, err := os.Open("db_config.json")
	if err != nil {
		return Config{}
	}
	defer file.Close()

	// 設定ファイルを読み込む
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return Config{}
	}

	// 設定ファイルをパースする
	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return Config{}
	}

	return config
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

// DB接続
func sqlConnect(config Config) (database *gorm.DB, err error) {
	DBMS := config.DBMS
	USER := config.User
	PASS := config.Pass
	PROTOCOL := config.Server
	DBNAME := config.DBName

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

// APITablesHandler file_dbsの中身をJSON形式で返すハンドラ
func apiTablesHandler(w http.ResponseWriter, r *http.Request) {
	// CORSのアクセス制御を行う
	w.Header().Set("Access-Control-Allow-Origin", "*")    // 任意のドメインからのアクセスを許可する
	w.Header().Set("Access-Control-Allow-Methods", "GET") // GETメソッドのみを許可する

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

	fileID, _ := strconv.Atoi(r.URL.Query().Get("fileId"))
	// 削除対象のファイルを特定し、削除する処理を実行する
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
	fileId, err := strconv.Atoi(r.URL.Query().Get("fileId"))
	fmt.Println("クリックしたファイルのID:", fileId)

	// ファイルのURLを取得する
	fileUrl, err := getFileUrl(fileId)
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
func getFileUrl(fileId int) (string, error) {
	// ファイルのURLを格納する変数
	var fileUrl string

	// file_dbsから指定したファイル名のレコードを1件取得する
	var fileDb File_dbs
	if err := db.Where("id = ?", fileId).First(&fileDb).Error; err != nil {
		fmt.Println("エラー(getFileUrl):", err)
		return "", err
	}

	// 取得したレコードのfilepathをfileUrlに代入する
	fileUrl = fileDb.Filepath

	return fileUrl, nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginHandlerが呼び出されました")
	w.Header().Set("Access-Control-Allow-Origin", "*")             // 任意のドメインからのアクセスを許可する
	w.Header().Set("Access-Control-Allow-Methods", "POST")         // POSTメソッドのみを許可する
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Content-Typeヘッダーのみを許可する
	username := r.FormValue("username")
	password := r.FormValue("password")

	// ユーザー名とパスワードを使用して、ログイン処理を行う
	var user UsersLoginfo
	if err := db.Where("username = ? and password = ?", username, password).First(&user).Error; err != nil {
		// ログインに失敗した場合
		if gorm.IsRecordNotFoundError(err) {
			// ユーザー名またはパスワードが間違っている
			resp := struct {
				Success bool   `json:"success"`
				Message string `json:"message"`
			}{
				Success: false,
				Message: "ユーザー名またはパスワードが間違っています",
			}
			respBytes, _ := json.Marshal(resp)
			w.Write(respBytes)
			return
		}
	}
	// ログインに成功した場合
	resp := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: true,
		Message: "ログインに成功しました",
	}
	respBytes, _ := json.Marshal(resp)
	w.Write(respBytes)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signupHandlerが呼び出されました")
	w.Header().Set("Access-Control-Allow-Origin", "*")             // 任意のドメインからのアクセスを許可する
	w.Header().Set("Access-Control-Allow-Methods", "POST")         // POSTメソッドのみを許可する
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Content-Typeヘッダーのみを許可する
	username := r.FormValue("username")
	password := r.FormValue("password")

	// ユーザー名とパスワードを使用して、サインアップ処理を行う
	//ユーザー名とパスワードをDBに保存する
	if err := db.Create(&UsersLoginfo{Username: username, Password: password}).Error; err != nil {
		resp := struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}{
			Success: false,
			Message: "既に登録されているユーザー名です",
		}
		respBytes, _ := json.Marshal(resp)
		w.Write(respBytes)
		return
	}

	resp := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: true,
		Message: "サインアップに成功しました",
	}
	respBytes, _ := json.Marshal(resp)
	w.Write(respBytes)
}

func setupRoutes(config Config) {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload/file", uploadHandler)
	mux.HandleFunc("/api/delete", deleteFileHandler)
	mux.HandleFunc("/api/tables", apiTablesHandler)
	mux.HandleFunc("/api/preview", apiPreviewHandler)
	mux.HandleFunc("/api/login", loginHandler)
	mux.HandleFunc("/api/signup", signupHandler)
	mux.Handle("/uploadfiles/", http.StripPrefix("/uploadfiles/", http.FileServer(http.Dir("./uploadfiles"))))

	if err := http.ListenAndServe(config.GoPort, mux); err != nil {
		log.Fatal(err)
	}
}

func serveVueApp() {
	fs := http.FileServer(http.Dir("./fronted/dist"))
	fmt.Println("Vue.jsアプリケーションをサーブ完了")
	http.Handle("/", fs)
}

func main() {
	// 設定値を取得する
	config := getConfig()
	// データベースと接続
	var err error
	db, err = sqlConnect(config)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	} else {
		fmt.Println("DB接続成功")
	}
	defer db.Close()

	setupRoutes(config)
	serveVueApp()
}
