<template>
  <div>
    <img alt="MinamiLab logo" src="./assets/minami_lab_logo.png">
    <h1>File Management</h1>
    <form @submit.prevent="uploadFile">
      <input type="file" ref="fileInput" />
      <button type="submit">Upload</button>
    </form>
    <p>filename:{{ file_name }}</p>
    <p>status:{{ upload_status }}</p>
    <h2>アップロード済みファイル一覧</h2>
    <!-- テーブル一覧を1つずつ表示する -->
    <table border style="margin: 0 auto" @click="showFile">
      <tr v-for="table in tables" :key="table">
        <td>{{ table }}</td>
      </tr>
    </table>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      file_name: null, //Goから受け取るアップロードされたファイル名
      upload_status: null, //Goから受け取るアップロードステータス
      tables: [], // Goから受け取るテーブル一覧を格納するデータプロパティ
    };
  },
  created() {
    // MySQLに接続してテーブル一覧を取得する
    this.getDB()
  },
  methods: {
    async uploadFile() {
      const file = this.$refs.fileInput.files[0];
      const formData = new FormData();
      formData.append('file', file);

      try {
        const response = await axios.post('http://localhost:12345/upload/file', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        },
          proxy: false //ローカルホストなのでプロキシを経由しない
        })
        if (response.data) { // レスポンスボディが存在する場合
          console.log("response.data", response.data)
          console.log(typeof response.data)
          this.file_name = response.data.filename // レスポンスボディのfilenameを参照して、responseMessageに代入する
          this.upload_status = response.data.status
          this.getDB()
        } else {
          console.error('レスポンスボディが存在しません')
        }
      } catch (error) {
        console.log('File uploaded Failed');
        console.error(error);
      }
    },
    getDB: function(){
      axios.get("http://localhost:12345/api/tables").then((res) => {
        this.tables = res.data;
      });
    },
    // tableをクリックした際に実行される処理
    async showFile(event) {
      console.log("click table!!!")
      // table要素がクリックされたときに、クリックされたtable要素のテキストを取得する
      var fileContent = event.target.textContent;
      fileContent = fileContent.replace(/\s/g, "");  // 文字列中のスペースを取り除く
      fileContent = fileContent.replace(/"/g, "");  // 文字列中の"を取り除く
      fileContent = fileContent.replace(/'/g, "");  // 文字列中の'を取り除く
      fileContent = fileContent.replace(/\n/g, "");  // 文字列中の改行を取り除く
      var fileNameElements = fileContent.split(",");  // 文字列を","で分割する
      let targetFilepath = "";  // 目的の文字列を格納する変数

      // 分割された要素を検索する
      for (const element of fileNameElements) {
        if (element.startsWith("filepath:")) {  // 要素が"filepath": の文字列で始まるかを確認する
          // "filepath": のあとの文字列を取り出す
          targetFilepath = element.replace("filepath:", "").trim();
          // targetFilepath = targetFilepath.replace('./', '');
          break;  // ループを抜ける
        }
      }
      console.log("targetFilepath", targetFilepath);  // 目的の文字列を出力する
      // サーバーに保管されているファイルをプレビューする
      axios.get(`http://localhost:12345/api/preview?fileName=${targetFilepath}`).then(response => {
          // プレビューするファイルのURLを取得する
          const fileUrl = response.data.fileUrl;
          // プレビューするファイルのURLをもとに、新しいタブを開く
          console.log("response.data", response.data)
          console.log("response.data", typeof response.data)
          console.log("fileUrl", fileUrl)
          window.open(fileUrl, '_blank');
        })
        .catch(error => {
          console.log("ファイルプレビューAPIでエラーが発生しました")
          console.error(error)
        });
    },
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
/* ヘッダーを装飾する */
header {
  background-color: #333;
  color: #fff;
  padding: 20px;
  text-align: center;
}

/* アップロードフォームを装飾する */
form {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

/* アップロードボタンを装飾する */
button[type="submit"] {
  background-color: #4caf50;
  border: none;
  color: white;
  padding: 15px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  cursor: pointer;
}

/* アップロード済みファイル一覧を装飾する */
table {
  border-collapse: collapse;
  width: 60%;
  margin: 0 auto;
}

td {
  border: 1px solid #ddd;
  padding: 8px;
}

tr:nth-child(even) {
  background-color: #f2f2f2;
}

tr:hover {
  background-color: #ddd;
}

th {
  padding-top: 12px;
  padding-bottom: 12px;
  text-align: left;
  background-color: #4caf50;
  color: white;
}
/* filename: 及び status: を装飾する */
p {
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: bold;
}

/* filename: の装飾をする */
p:nth-child(1) {
  color: blue;
}

/* status: の装飾をする */
p:nth-child(2) {
  color: green;
}
</style>
