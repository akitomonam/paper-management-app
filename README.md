# Paper Management Web App
論文を管理するWEBアプリケーションです
# DEMO
![demo2](sample_demo/demo-v2.gif)
# Features
- pdf, pptxファイルのアップロード
- 検索
- 並べ替え
- プレビュー
- 削除
- コメント
- メタ情報編集(手動 or 自動)
- 論文自動翻訳
- 論文キーワード
- お気に入り
- マイページ
- ユーザー登録
- ログイン
- ログアウト
# Requirement
* Docker
* Docker Compose
# Preparations
以下のファイルの中身を修正性.sampleを取り除いたファイルを作成
- [.sample.env](.sample.env)
- [config.sample.json](server/backend_python/config.sample.json)
- [db_config.sample.json](server/backend_go/db_config.sample.json)
- [config.sample.js](client/Vue3/config.sample.js)
# Usage
以下のコマンドでアプリケーションを起動
```
docker-compose up -d
```
以下のアドレスにブラウザでアクセス
```
http://localhost:8080/
```
