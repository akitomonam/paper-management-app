# 論文管理 WEB アプリ

論文を管理する WEB アプリです。  
研究室運営や自然言語処理技術向上を目的として作成しました。  
![demo](https://user-images.githubusercontent.com/72239675/210528954-489326a6-5e1d-4612-8b6f-cf472112d200.gif)

# 仕様

- ファイルをアップロードし、それをサーバー側で DB(mySQL)に保存
- DB(mySQL)を参照してアップロード済みファイル一覧を表示
- アップロード済みファイルを、新しいタブでプレビュー
- アップロード済みファイルの削除
- マイページ機能
- アップロード済みファイルにコメント
- お気に入り機能

# 実行

```
docker-compose up -d
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
