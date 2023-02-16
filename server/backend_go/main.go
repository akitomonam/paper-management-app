// main.go
package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"time"
	"upload/file"

	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
)

type Config struct {
	DBMS   string `json:"dbms"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Server string `json:"server"`
	DBName string `json:"dbname"`
	GoPort string `json:"go_port"`
}

type Users struct {
	ID        int    `db:"id"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	File_path string `db:"file_path"`
}

// Papers ファイルのメタ情報
type Papers struct {
	ID         int
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Publisher  string    `json:"publisher"`
	Year       int       `json:"year"`
	Abstract   string    `json:"abstract"`
	File_name  string    `json:"file_name"`
	File_path  string    `json:"file_path"`
	User_id    int       `json:"user_id"`
	Created_at time.Time `gorm:"not null"`
}

type supportFiles struct {
	ID         int
	PaperId    int       `json:"paperId"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Publisher  string    `json:"publisher"`
	Year       int       `json:"year"`
	Abstract   string    `json:"abstract"`
	File_name  string    `json:"file_name"`
	File_path  string    `json:"file_path"`
	User_id    int       `json:"user_id"`
	Created_at time.Time `gorm:"not null"`
}

type Comment struct {
	ID        int       `gorm:"primary_key;auto_increment"`
	PaperID   int       `gorm:"not null"`
	UserID    int       `gorm:"not null"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"not null"`
}

type Favorites struct {
	ID       int `db:"id"`
	Paper_id int `db:"paper_id"`
	User_id  int `db:"user_id"`
	Rating   int `db:"rating"`
}

type Keywords struct {
	ID       int    `db:"id"`
	Paper_id int    `db:"paper_id"`
	Keyword  string `db:"keyword"`
	Flag     bool   `db:"flag"`
}

type Sessions struct {
	Session_Token string `db:"session_token"`
	User_id       int    `db:"user_id"`
}

var db *gorm.DB
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

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

// APITablesHandler Papersの中身をJSON形式で返すハンドラ
func apiTablesHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken := r.URL.Query().Get("sessionToken")
	favorite_flag := r.URL.Query().Get("favorite")
	// sessionToken, _ = url.QueryUnescape(sessionToken)
	fmt.Println("table-list-sessionToken:", sessionToken)
	fmt.Println("favorite_flag:", favorite_flag)

	var fileDbs []Papers
	if sessionToken == "" {
		fmt.Println("セッショントークンなし")
		// Papersからすべてのレコードを取得する
		if err := db.Find(&fileDbs).Error; err != nil {
			// エラーを出力する
			fmt.Println("エラー:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		fmt.Println("セッショントークンあり")
		session, err := FindSession(sessionToken)
		if err != nil {
			fmt.Println("セッショントークンが無効です。:")
			http.Error(w, "セッショントークンが無効です。", http.StatusUnauthorized)
			return
		}

		user_id := session.User_id

		if favorite_flag == "true" {
			// user_idを使用して、favoritesテーブルのratingが1以上であるpapers_idを用いてPapersからレコードを取得する
			var favoriteIDs []int
			if err := db.Table("favorites").Where("user_id = ? AND rating >= 1", user_id).Pluck("paper_id", &favoriteIDs).Error; err != nil {
				fmt.Println("エラー:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := db.Where("ID IN (?)", favoriteIDs).Find(&fileDbs).Error; err != nil {
				fmt.Println("エラー:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			//user_idを使用して、Papersからレコードを取得する
			if err := db.Where("user_id = ?", user_id).Find(&fileDbs).Error; err != nil {
				// エラーを出力する
				fmt.Println("エラー:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	// Papersの中身をJSON形式で返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(fileDbs); err != nil {
		fmt.Println("エラー:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func apiSupportFilesHandler(w http.ResponseWriter, r *http.Request) {
	paperId := r.URL.Query().Get("paperId")
	fmt.Println("paperId:", paperId)

	var fileDbs []supportFiles
	//paperIdを使用して、supportFilesからレコードを取得する
	if err := db.Where("paper_id = ?", paperId).Find(&fileDbs).Error; err != nil {
		// エラーを出力する
		fmt.Println("エラー:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Papersの中身をJSON形式で返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(fileDbs); err != nil {
		fmt.Println("エラー:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteFileHandlerが呼び出されました")

	fileID, _ := strconv.Atoi(r.URL.Query().Get("fileId"))
	// 削除対象のファイルを特定し、削除する処理を実行する
	// PapersテーブルからIDを指定してレコードを取得する
	var fileDb Papers
	if err := db.Where("id = ?", fileID).First(&fileDb).Error; err != nil {
		// http.Error(w, "レコードが見つかりません", http.StatusNotFound)
		fmt.Println("レコードが見つかりません")
		fmt.Println("レコードが見つかりません:", err)
		// レスポンスを返す
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"result": "false"})
		return
	}
	// Papersテーブルからレコードを削除する
	fmt.Println("filedb:", &fileDb)
	if err := db.Delete(&fileDb).Error; err != nil {
		http.Error(w, "データの削除に失敗しました", http.StatusInternalServerError)
		fmt.Println("データの削除に失敗しました")
		return
	}
	//DBから取得したファイルパスでファイルの削除
	if err := os.Remove(fileDb.File_path); err != nil {
		http.Error(w, "ファイルの削除に失敗しました", http.StatusInternalServerError)
		fmt.Println("ファイルの削除に失敗しました")
		return
	}
	// レスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": "true"})
	fmt.Println("finish deletefile")
}

func deleteSupportFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteFileHandlerが呼び出されました")

	fileID, _ := strconv.Atoi(r.URL.Query().Get("fileId"))
	// 削除対象のファイルを特定し、削除する処理を実行する
	// PapersテーブルからIDを指定してレコードを取得する
	var fileDb supportFiles
	if err := db.Where("id = ?", fileID).First(&fileDb).Error; err != nil {
		// http.Error(w, "レコードが見つかりません", http.StatusNotFound)
		fmt.Println("レコードが見つかりません")
		fmt.Println("レコードが見つかりません:", err)
		// レスポンスを返す
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"result": "false"})
		return
	}
	// Papersテーブルからレコードを削除する
	fmt.Println("filedb:", &fileDb)
	if err := db.Delete(&fileDb).Error; err != nil {
		http.Error(w, "データの削除に失敗しました", http.StatusInternalServerError)
		fmt.Println("データの削除に失敗しました")
		return
	}
	//DBから取得したファイルパスでファイルの削除
	if err := os.Remove(fileDb.File_path); err != nil {
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
	// ファイル名を取得する
	fileId, err := strconv.Atoi(r.URL.Query().Get("fileId"))
	fmt.Println("クリックしたファイルのID:", fileId)

	// ファイルのURLを取得する
	fileUrl, err := getFileUrl(fileId)
	//適したURLの形にする
	// fileUrl = fileUrl[1:]
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

func apiPreviewSupportFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("apiPreviewHandlerが呼び出されました")
	// ファイル名を取得する
	fileId, err := strconv.Atoi(r.URL.Query().Get("fileId"))
	fmt.Println("クリックしたファイルのID:", fileId)

	// ファイルのURLを取得する
	fileUrl, err := getSupportFileUrl(fileId)
	//適したURLの形にする
	// fileUrl = fileUrl[1:]
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

	// Papersから指定したファイル名のレコードを1件取得する
	var fileDb Papers
	if err := db.Where("id = ?", fileId).First(&fileDb).Error; err != nil {
		fmt.Println("エラー(getFileUrl):", err)
		return "", err
	}

	// 取得したレコードのfilepathをfileUrlに代入する
	fileUrl = fileDb.File_path

	return fileUrl, nil
}

// ファイルのURLを取得する関数
func getSupportFileUrl(fileId int) (string, error) {
	// ファイルのURLを格納する変数
	var fileUrl string

	// Papersから指定したファイル名のレコードを1件取得する
	var fileDb supportFiles
	if err := db.Where("id = ?", fileId).First(&fileDb).Error; err != nil {
		fmt.Println("エラー(getFileUrl):", err)
		return "", err
	}

	// 取得したレコードのfilepathをfileUrlに代入する
	fileUrl = fileDb.File_path

	return fileUrl, nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginHandlerが呼び出されました")
	username := r.FormValue("username")
	password := r.FormValue("password")

	// ユーザー名とパスワードを使用して、ログイン処理を行う
	var user Users
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
	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println("セッションGET失敗")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// トークンを生成する
	b := make([]byte, 32)
	rand.Read(b)
	token := base64.StdEncoding.EncodeToString(b)
	token = removeSpecialCharacters(token)
	// セッションにトークンを保存する(必要ないかも)
	session.Values["token"] = token
	session.Save(r, w)

	user_id := user.ID

	if err := db.Create(&Sessions{User_id: user_id, Session_Token: token}).Error; err != nil {
		resp := struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}{
			Success: false,
			Message: "セッション取得に失敗しました。",
		}
		respBytes, _ := json.Marshal(resp)
		w.Write(respBytes)
		return
	}

	resp := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Token   string `json:"token"`
	}{
		Success: true,
		Message: "ログインに成功しました",
		Token:   token,
	}
	respBytes, _ := json.Marshal(resp)
	w.Write(respBytes)
}

func removeSpecialCharacters(s string) string {
	// +や-を除去するための正規表現
	re := regexp.MustCompile(`[+-]`)
	return re.ReplaceAllString(s, "")
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signupHandlerが呼び出されました")
	username := r.FormValue("username")
	password := r.FormValue("password")

	// ユーザー名とパスワードを使用して、サインアップ処理を行う
	//ユーザー名とパスワードをDBに保存する
	if err := db.Create(&Users{Username: username, Password: password}).Error; err != nil {
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

func userinfoHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken := r.URL.Query().Get("sessionToken")
	fmt.Println("sessionToken:", sessionToken)

	session, err := FindSession(sessionToken)
	if err != nil {
		fmt.Println("セッショントークンが無効です。:")
		http.Error(w, "セッショントークンが無効です。", http.StatusUnauthorized)
		return
	}

	user_id := session.User_id

	var userDb Users
	if err := db.Where("id = ?", user_id).First(&userDb).Error; err != nil {
		fmt.Println("エラー(useInfo):", err)
		return
	}

	// ユーザー情報を必要な情報のみ返す
	userInfo := Users{
		Username:  userDb.Username,
		File_path: userDb.File_path,
	}

	res, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println("エラー(useInfo):", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func userlistHandler(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Username string
		Filepath string
	}

	var UserList []User
	// Userからすべてのレコードを取得する
	if err := db.Find(&UserList).Error; err != nil {
		// エラーを出力する
		fmt.Println("エラー:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Userの中身をJSON形式で返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(UserList); err != nil {
		fmt.Println("エラー:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func FindSession(sessionToken string) (Sessions, error) {
	// セッションDBを検索
	var session Sessions
	err := db.Where("session_token = ?", sessionToken).First(&session).Error
	if err != nil {
		// セッションが見つからない場合や、エラーが発生した場合
		return Sessions{}, fmt.Errorf("セッションが無効です: %v", err)
	}
	return session, nil
}

func apiPapersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// URL パスからパラメータを取得する
	// URL パスの最後のセグメントを取得する
	_, id := path.Split(r.URL.Path)

	// id を使って、DB から論文の詳細情報を取得する処理を記述する
	var Paper Papers
	if err := db.Where("id = ?", id).First(&Paper).Error; err != nil {
		// http.Error(w, "レコードが見つかりません", http.StatusNotFound)
		fmt.Println("レコードが見つかりません")
		fmt.Println("レコードが見つかりません:", err)
		// レスポンスを返す
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"result": "false"})
		return
	}
	//keywordsデータベースから該当するidでflagがtrueのものを取得する
	var Keyword []Keywords
	if err := db.Where("paper_id = ? AND flag = ?", id, true).Find(&Keyword).Error; err != nil {
		fmt.Println("レコードが見つかりません")
		fmt.Println("レコードが見つかりません:", err)
		// レスポンスを返す
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"result": "false"})
		return
	}

	// 論文情報とkeywordをまとめて返す
	type PaperInfo struct {
		Paper    Papers
		Keywords []Keywords
	}
	paperInfo := PaperInfo{
		Paper:    Paper,
		Keywords: Keyword,
	}

	res, err := json.Marshal(paperInfo)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func apiEditPaperInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("apiEditPaperInfoHandlerが呼び出されました")
	// HTTPメソッドがOPTIONSの場合は、ここで処理を終了する
	if r.Method == http.MethodOptions {
		return
	}

	//To allocate slice for request body
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//parse json
	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id := int(jsonBody["id"].(float64))
	title := jsonBody["title"].(string)
	abstract := jsonBody["abstract"].(string)
	author := jsonBody["author"].(string)
	publisher := jsonBody["publisher"].(string)
	year := int(jsonBody["year"].(float64))

	// id を使って、DB から論文の詳細情報を取得して、「Title,Author,Publisher,Year,Abstract」を更新する
	var paper Papers
	if err := db.Where("id = ?", id).First(&paper).Error; err != nil {
		http.Error(w, "Paper not found", http.StatusNotFound)
		return
	}
	paper.Title = title
	paper.Abstract = abstract
	paper.Author = author
	paper.Publisher = publisher
	paper.Year = year
	if err := db.Save(&paper).Error; err != nil {
		http.Error(w, "Error updating paper", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(paper)
	if err != nil {
		fmt.Println("エラー(paper):", err)
		return
	}

	for _, keyword := range jsonBody["keywords"].([]interface{}) {
		kw, _ := keyword.(map[string]interface{})
		fmt.Println(kw["Keyword"], ":", kw["Flag"])
		fmt.Printf("flag:%T\n", kw["Flag"])
		//DBに反映する(DBに同じキーワードと同じidが存在しない場合は新規追加する、存在していた場合はflagを更新する)
		var keyword Keywords
		if err := db.Where("keyword = ? AND paper_id = ?", kw["Keyword"], id).First(&keyword).Error; err != nil {
			//新規追加
			db.Create(&Keywords{Keyword: kw["Keyword"].(string), Paper_id: id, Flag: kw["Flag"].(bool)})
		} else {
			//更新
			keyword.Flag = kw["Flag"].(bool)
			db.Save(&keyword)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func apiFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("apiFavoriteHandlerが呼び出されました")

	// HTTPメソッドがOPTIONSの場合は、ここで処理を終了する
	if r.Method == http.MethodOptions {
		return
	}

	paper_id, err := strconv.Atoi(r.FormValue("paperId"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	session_token := r.FormValue("sessionToken")
	rating, err := strconv.Atoi(r.FormValue("rating"))
	if err != nil {
		http.Error(w, "Invalid Year", http.StatusBadRequest)
		return
	}

	//session_tokenからuser_idを取得
	var session Sessions
	db.Where("session_token = ?", session_token).First(&session)
	if session.User_id == 0 {
		http.Error(w, "Session Token is not found", http.StatusUnauthorized)
		return
	}
	user_id := session.User_id

	// favoritesテーブルに既にpaper_id、user_idと同じレコードがあった場合、そのレコードのratingを更新
	var favorite Favorites
	db.Where("paper_id = ? AND user_id = ?", paper_id, user_id).First(&favorite)
	if favorite.ID != 0 {
		favorite.Rating = rating
		db.Save(&favorite)
		return
	}

	// なかった場合、favoriteレコードを新規作成
	favorite = Favorites{Paper_id: paper_id, User_id: user_id, Rating: rating}
	db.Create(&favorite)

	resp := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: true,
		Message: "お気に入り登録に成功しました",
	}
	respBytes, _ := json.Marshal(resp)
	w.Write(respBytes)
}

func apicheckFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("apiFavoriteHandlerが呼び出されました")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// クエリパラメータからセッショントークンとペーパーIDを取得する
	session_token := r.URL.Query().Get("sessionToken")
	paperIdStr := r.URL.Query().Get("paperId")
	paper_id, _ := strconv.Atoi(paperIdStr)

	//session_tokenからuser_idを取得
	var session Sessions
	db.Where("session_token = ?", session_token).First(&session)
	if session.User_id == 0 {
		http.Error(w, "Session Token is not found", http.StatusUnauthorized)
		return
	}
	user_id := session.User_id

	// favoritesテーブルからpaper_id、user_idを用いてratingを取得
	var favorite Favorites
	db.Where("paper_id = ? AND user_id = ?", paper_id, user_id).First(&favorite)
	if favorite.ID == 0 {
		favorite.Rating = 0
	}

	// favoriteレコードが見つかった場合は、ratingを返す
	resp := struct {
		Rating  int    `json:"rating"`
		Message string `json:"message"`
	}{
		Rating:  favorite.Rating,
		Message: "お気に入り登録確認成功",
	}
	respBytes, _ := json.Marshal(resp)
	w.Write(respBytes)
}

func apiCommentPreviewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("apiCommentPreviewHandlerが呼び出されました")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// クエリパラメータからペーパーIDを取得する
	paperIdStr := r.URL.Query().Get("paperId")
	paper_id, _ := strconv.Atoi(paperIdStr)
	fmt.Println("paperIdStr", paperIdStr)

	var comments []Comment

	// paper_idを用いてcommentsからレコードを取得する
	if err := db.Where("paper_id = ?", paper_id).Find(&comments).Error; err != nil {
		// エラーを出力する
		fmt.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// commentsのuser_idを用いてusersのusernameを紐づける(マージする)
	type CommentPreview struct {
		ID        int       `gorm:"primary_key;auto_increment"`
		PaperID   int       `gorm:"not null"`
		UserID    int       `gorm:"not null"`
		Content   string    `gorm:"type:text;not null"`
		CreatedAt time.Time `gorm:"not null"`
		Username  string
	}

	var commentPreviews []CommentPreview

	var users []Users

	userIDs := []int{}
	for _, comment := range comments {
		userIDs = append(userIDs, comment.UserID)
	}

	// TODO:マージしたほうが速い

	if err := db.Select("id, username").Where("id IN (?)", userIDs).Find(&users).Error; err != nil {
		fmt.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userMap := make(map[int]string)
	for _, user := range users {
		userMap[user.ID] = user.Username
	}

	for _, comment := range comments {
		commentPreviews = append(commentPreviews, CommentPreview{
			ID:        comment.ID,
			PaperID:   comment.PaperID,
			UserID:    comment.UserID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			Username:  userMap[comment.UserID],
		})
	}

	// commentPreviewsの中身をJSON形式で返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(commentPreviews); err != nil {
		fmt.Println("エラー:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func apiCommentAddHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("apiCommentAddHandlerが呼び出されました")
	// HTTPメソッドがOPTIONSの場合は、ここで処理を終了する
	if r.Method == http.MethodOptions {
		return
	}

	// クエリパラメータからセッショントークンとペーパーIDを取得する
	session_token := r.FormValue("sessionToken")
	paperIdStr := r.FormValue("paperId")
	paper_id, _ := strconv.Atoi(paperIdStr)
	content := r.FormValue("comments")
	fmt.Println("session_token", session_token)
	fmt.Println("paperIdStr", paperIdStr)
	fmt.Println("content", content)

	//session_tokenからuser_idを取得
	var session Sessions
	db.Where("session_token = ?", session_token).First(&session)
	if session.User_id == 0 {
		http.Error(w, "Session Token is not found", http.StatusUnauthorized)
		return
	}
	user_id := session.User_id

	var comments []Comment

	// paper_id, user_idを使用して、commentsからレコードを取得する
	if err := db.Where("paper_id = ? and user_id = ?", paper_id, user_id).Find(&comments).Error; err != nil {
		// エラーを出力する
		fmt.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := db.Create(&Comment{Content: content, PaperID: paper_id, UserID: user_id}).Error; err != nil {
		resp := struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}{
			Success: false,
			Message: "コメント投稿でエラーが発生しました",
		}
		respBytes, _ := json.Marshal(resp)
		w.Write(respBytes)
		return
	}

	resp := struct {
		// Rating  int    `json:"rating"`
		Message string `json:"message"`
	}{
		// Rating:  favorite.Rating,
		Message: "コメント投稿成功",
	}
	respBytes, _ := json.Marshal(resp)
	w.Write(respBytes)
}

func setupRoutes(config Config) {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload/file", corsMiddleware(uploadHandler))
	mux.HandleFunc("/api/delete", corsMiddleware(deleteFileHandler))
	mux.HandleFunc("/api/deleteSupportFile", corsMiddleware(deleteSupportFileHandler))
	mux.HandleFunc("/api/tables", corsMiddleware(apiTablesHandler))
	mux.HandleFunc("/api/supportFiles", corsMiddleware(apiSupportFilesHandler))
	mux.HandleFunc("/api/preview", corsMiddleware(apiPreviewHandler))
	mux.HandleFunc("/api/previewSupportFile", corsMiddleware(apiPreviewSupportFileHandler))
	mux.HandleFunc("/api/login", corsMiddleware(loginHandler))
	mux.HandleFunc("/api/signup", corsMiddleware(signupHandler))
	mux.HandleFunc("/api/userinfo", corsMiddleware(userinfoHandler))
	mux.HandleFunc("/api/userlist", corsMiddleware(userlistHandler))
	mux.HandleFunc("/api/papers/", corsMiddleware(apiPapersHandler))
	mux.HandleFunc("/api/editpaperinfo", corsMiddleware(apiEditPaperInfoHandler))
	mux.HandleFunc("/api/favorite", corsMiddleware(apiFavoriteHandler))
	mux.HandleFunc("/api/checkFavorite", corsMiddleware(apicheckFavoriteHandler))
	mux.HandleFunc("/api/comment_preview", corsMiddleware(apiCommentPreviewHandler))
	mux.HandleFunc("/api/comment_add", corsMiddleware(apiCommentAddHandler))
	mux.Handle("/mnt/uploadfiles/", http.StripPrefix("/mnt/uploadfiles", http.FileServer(http.Dir("/mnt/uploadfiles"))))

	if err := http.ListenAndServe(config.GoPort, mux); err != nil {
		log.Fatal(err)
	}
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, paperId")
		next(w, r)
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
