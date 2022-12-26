# ファイルアップロードWEBアプリ
<img width="631" alt="image" src="https://user-images.githubusercontent.com/72239675/209383133-4c4008d9-71dc-4682-adec-683825b01ea4.png">

# 仕様
- ファイルをアップロードし、それをサーバー側でDBに保存。  
- DBを参照してアップロード済みファイル一覧を表示。
# 環境構築
Goをインストール
```
sudo apt install golang
go version
```
MySQLをインストール
```
sudo apt install mysql-server
sudo service mysql start
sudo service mysql status
```
Vueのインストール
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
