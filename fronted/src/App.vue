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
    <table border style="margin: 0 auto">
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
      file_name: 'None', //Goから受け取るアップロードされたファイル名
      upload_status: 'None', //Goから受け取るアップロードステータス
      tables: [], // Goから受け取るテーブル一覧を格納するデータプロパティ
    };
  },
  created() {
    // MySQLに接続してテーブル一覧を取得する
    axios.get("http://localhost:12345/api/tables").then((res) => {
      this.tables = res.data;
    });

  },
  methods: {
    async uploadFile() {
      const file = this.$refs.fileInput.files[0];
      const formData = new FormData();
      formData.append('file', file);

      try {
        const response = await axios.post('http://localhost:12345/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        },
          proxy: false //ローカルホストなのでプロキシを経由しない
        })
        if (response.data) { // レスポンスボディが存在する場合
          this.file_name = response.data.filename // レスポンスボディのfilenameを参照して、responseMessageに代入する
          this.upload_status = response.data.status
        } else {
          console.error('レスポンスボディが存在しません')
        }
      } catch (error) {
        console.log('File uploaded Failed');
        console.error(error);
      }
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
</style>
