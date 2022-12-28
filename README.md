# 論文管理 WEB アプリ

![image](https://user-images.githubusercontent.com/72239675/209828493-44812c03-32c5-42d7-884a-9789c88e2875.png)

# 仕様

- ファイルをアップロードし、それをサーバー側で DB に保存
- DB を参照してアップロード済みファイル一覧を表示
- アップロード済みファイル一覧をクリックすると、新しいタブでプレビュー
- アップロード済みファイルの削除
- アップロード済みファイルにコメント

# 環境構築

Go をインストール

```
sudo apt install golang
go version
```

MySQL をインストール

```
sudo apt install mysql-server
sudo service mysql start
sudo service mysql status
```

Vue のインストール

# 実行

```
cd ./fronted
npm run serve
cd ..
go run main.go
```

# 検証環境

```
Distributor ID: Ubuntu
Description:    Ubuntu 20.04.5 LTS
Release:        20.04
Codename:       focal
```

# 参考

- https://zenn.dev/khayakawa/articles/file-upload-go
- https://rightcode.co.jp/blog/information-technology/golang-introduction-mysql-connection
