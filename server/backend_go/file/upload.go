// upload.go
package file

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const MaxUploadSize = 100 * 1024 * 1024 // 最大ファイルサイズ

type Papers struct {
	ID         int
	Title      string `json:"title"`
	Author     string `json:"author"`
	Publisher  string `json:"publisher"`
	Year       int    `json:"year"`
	Abstract   string `json:"abstract"`
	File_name  string `json:"file_name"`
	File_path  string `json:"file_path"`
	User_id    int    `json:"user_id"`
	Created_at string `json:"created_at" sql:"not null;type:datetime"`
	// Createdat time.Time `json:"created_at" sql:"not null;type:datetime"`
}

type Sessions struct {
	SessionToken string `db:"session_token"`
	User_id      int    `db:"user_id"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	status := "True"
	// CORSのアクセス制御を行う
	w.Header().Set("Access-Control-Allow-Origin", "*")                            // 任意のドメインからのアクセスを許可する
	w.Header().Set("Access-Control-Allow-Methods", "POST")                        // POSTメソッドのみを許可する
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") //  Content-Type, Authorizationヘッダーのみを許可する

	// HTTPメソッドがOPTIONSの場合は、ここで処理を終了する
	if r.Method == http.MethodOptions {
		return
	}

	// セッショントークンを受け取る
	sessionToken := r.Header.Get("Authorization")
	sessionToken = strings.TrimSpace(strings.TrimPrefix(sessionToken, "Bearer "))
	fmt.Println("sessionToken:", sessionToken)
	// セッションDBを検索
	session, err := FindSession(sessionToken)
	if err != nil {
		// セッショントークンが無効の場合
		status = "False"
		fmt.Println("セッショントークンが無効です。:")
		http.Error(w, "セッショントークンが無効です。", http.StatusUnauthorized)
		return
	}

	// 有効である場合は、ユーザーIDを取得する
	user_id := session.User_id

	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		status = "False"
		http.Error(w, "1MB以下のファイルを選択してください。", http.StatusBadRequest)
		fmt.Println("ParseMultipartFormでエラーが発生:", err)
		return
	}

	// フォームのファイルを取得する
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		status = "False"
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("ParseMultipartFormでエラーが発生:", err)
		return
	}
	defer file.Close()

	// 保存用のディレクトリを作成する（存在していなければ、保存用のディレクトリを新規作成）
	err = os.MkdirAll("/mnt/uploadfiles", os.ModePerm)
	if err != nil {
		status = "False"
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// アップロードされたファイルをバッファに読み込む
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		status = "False"
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filetype := http.DetectContentType(buff)
	fmt.Println("filetype:", filetype)
	if filetype != "image/jpeg" && filetype != "image/png" && filetype != "application/pdf" && filetype != "application/zip" {
		status = "False"
		http.Error(w, "JPEG、PNG、PDF、または、パワーポイントでアップロードしてください。", http.StatusBadRequest)
		return
	}

	// ファイルを読み込み開始位置に戻す
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ファイル名を取得する
	filename := fileHeader.Filename
	// ファイル名から拡張子を取り出す
	ext := filepath.Ext(filename)

	// 保存先のパスを生成する
	savePath := "/mnt/uploadfiles/" + filename

	// os.Statで、保存先のパスが存在するか確認する
	_, err = os.Stat(savePath)
	if err == nil {
		// 同名のファイルがある場合は、番号をつけて保存するようにする
		// 番号をつける
		i := 1
		for {
			savePath = "/mnt/uploadfiles/" + filename[:len(filename)-len(ext)] + "(" + fmt.Sprint(i) + ")" + ext
			_, err = os.Stat(savePath)
			if err != nil {
				break
			}
			i++
		}
		// ファイル名を更新する
		filename = filename[:len(filename)-len(ext)] + "(" + fmt.Sprint(i) + ")" + ext
	}

	// 保存するファイルを作成する
	dst, err := os.Create(savePath)
	if err != nil {
		status = "False"
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// アップロードしたファイルを保存用のディレクトリにコピーする
	_, err = io.Copy(dst, file)
	if err != nil {
		status = "False"
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//DBにパスなどのメタ情報を保存
	add2sql(filename, savePath, user_id)

	// アップロード成功のレスポンスを返す
	response := map[string]string{
		"status":   status,
		"filename": filename,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "sdkogaken"
	PROTOCOL := "tcp(db-mysql)"
	DBNAME := "test_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

func add2sql(file_name string, file_path string, user_id int) {
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功(ファイルアップロード)")
	}
	defer db.Close()

	error := db.Create(&Papers{
		File_name:  file_name,
		File_path:  file_path,
		User_id:    user_id,
		Created_at: getDate(),
	}).Error
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("データ追加成功")
	}
}

func getDate() string {
	const layout = "2006-01-02 15:04:05"
	now := time.Now()
	return now.Format(layout)
}

func FindSession(sessionToken string) (Sessions, error) {
	// セッションDBを検索
	var session Sessions
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功(セッション検索)")
	}
	defer db.Close()
	err = db.Where("session_token = ?", sessionToken).First(&session).Error
	if err != nil {
		// セッションが見つからない場合や、エラーが発生した場合
		return Sessions{}, fmt.Errorf("セッションが無効です: %v", err)
	}
	return session, nil
}
