// upload.go
package file

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const MaxUploadSize = 1024 * 1024 // 最大ファイルサイズ

// File_dbs ファイルのメタ情報
type File_dbs struct {
	ID       int
	Filename string `json:"filename"`
	Filepath string `json:"filepath"`
	Updateat string `json:"updateAt" sql:"not null;type:datetime"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// CORSのアクセス制御を行う
	w.Header().Set("Access-Control-Allow-Origin", "*")             // 任意のドメインからのアクセスを許可する
	w.Header().Set("Access-Control-Allow-Methods", "POST")         // POSTメソッドのみを許可する
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Content-Typeヘッダーのみを許可する

	if r.Method != "POST" {
		http.Error(w, "処理を終了します。", http.StatusMethodNotAllowed)
		return
	}

	// fmt.Println("POSTリクエストを受け取りました")
	// fmt.Println("リクエストのURL:", r.URL)
	// fmt.Println("リクエストのヘッダー:", r.Header)
	// fmt.Println("リクエストのボディ:", r.Body)

	// return

	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		http.Error(w, "1MB以下のファイルを選択してください。", http.StatusBadRequest)
		fmt.Println("ParseMultipartFormでエラーが発生:", err)
		return
	}

	// フォームのファイルを取得する
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("ParseMultipartFormでエラーが発生:", err)
		return
	}
	defer file.Close()

	// 保存用のディレクトリを作成する（存在していなければ、保存用のディレクトリを新規作成）
	err = os.MkdirAll("./uploadfiles", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// アップロードされたファイルをバッファに読み込む
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" && filetype != "application/pdf" {
		http.Error(w, "JPEG、PNG、または、PDFでアップロードしてください。", http.StatusBadRequest)
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
	savePath := "./uploadfiles/" + filename

	// os.Statで、保存先のパスが存在するか確認する
	_, err = os.Stat(savePath)
	if err == nil {
		// 同名のファイルがある場合は、番号をつけて保存するようにする
		// 番号をつける
		i := 1
		for {
			savePath = "./uploadfiles/" + filename[:len(filename)-len(ext)] + "(" + fmt.Sprint(i) + ")" + ext
			_, err = os.Stat(savePath)
			if err != nil {
				break
			}
			i++
		}
	}

	// 保存するファイルを作成する
	dst, err := os.Create(savePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// アップロードしたファイルを保存用のディレクトリにコピーする
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//DBにパスなどのメタ情報を保存
	add2sql(filename, savePath)

	// アップロード成功のレスポンスを返す
	response := map[string]string{
		"status":   "success",
		"filename": fileHeader.Filename,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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

func add2sql(file_name string, file_path string) {
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	defer db.Close()

	error := db.Create(&File_dbs{
		Filename: file_name,
		Filepath: file_path,
		Updateat: getDate(),
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
