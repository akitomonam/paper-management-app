<template>
  <div>
    <header style="display: flex; align-items: center; justify-content: center; background-color: transparent">
      <img alt="MinamiLab logo" src="./assets/minami_lab_logo.png" style="height: 100px; width: 100px">
      <h1 style="color: #000">Paper Management Site</h1>
    </header>
    <form @submit.prevent="uploadFile">
      <input type="file" ref="fileInput" />
      <button type="submit">Upload</button>
    </form>
    <p>filename:{{ file_name }}</p>
    <p>status:{{ upload_status }}</p>
    <h2>Uploaded File List</h2>
    <!-- テーブル一覧を1つずつ表示する -->
    <table>
      <div class="draggable-container" style="display: flex; justify-content: center">
      <draggable v-model="tables" group="people" item-key="ID" handle=".handle">
          <template #item="{element}">
            <tr style="border: solid 1px #000;">
                <div class="drag-item">
                  <span class="handle">・</span>
                  {{ element.filename }}
                  <button @click="showFile((element.filepath))">Preview</button>
                </div>
              </tr>
          </template>
      </draggable>
    </div>
    </table>
    <footer>
      <p>Created:Akitomo SATO</p>
      <p>University of Electro-communications</p>
    </footer>
  </div>
</template>

<script>
import axios from 'axios';
import draggable from 'vuedraggable';
import { config } from '../config';

export default {
  components: {
    draggable,
  },
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
        const response = await axios.post(`${config.URL}:${config.PORT}/upload/file`, formData, {
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
      axios.get(`${config.URL}:${config.PORT}/api/tables`).then((res) => {
        this.tables = res.data;
      });
    },
    // tableをクリックした際に実行される処理
    async showFile(targetFilepath) {
      console.log("click table!!!")
      console.log("targetFilepath", targetFilepath);  // 目的の文字列を出力する
      // サーバーに保管されているファイルをプレビューする
      axios.get(`${config.URL}:${config.PORT}/api/preview?fileName=${targetFilepath}`).then(response => {
          // プレビューするファイルのURLを取得する
          let filepath = response.data.fileUrl;
          // プレビューするファイルのURLをもとに、新しいタブを開く
          filepath = filepath.replace(/^\./, '');
          window.open(`${config.URL}:${config.PORT}${filepath}`, '_blank');
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

footer {
  background-color: rgba(255, 255, 255, 0);
  padding: 1px;
  font-size: 1px;
}
</style>
