// upload.go
package file

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const MaxUploadSize = 1024 * 1024 // 最大ファイルサイズ

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "処理を終了します。", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		http.Error(w, "1MB以下のファイルを選択してください。", http.StatusBadRequest)
	}

	// 1. フォームのファイルを取得する
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 2. 保存用のディレクトリを作成する（存在していなければ、保存用のディレクトリを新規作成）
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

	// 3. 保存するファイルを作成する
	dst, err := os.Create(fmt.Sprintf("./uploadfiles/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// 4. アップロードしたファイルを保存用のディレクトリにコピーする
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "アップロード成功！")
}
