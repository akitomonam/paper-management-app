<template>
  <div>
    <HeaderComponent />
    <UploadForm @update-upload-status="updateUploadStatus" />
    <!-- <p>filename:{{ file_name }}</p>
    <p>status:{{ upload_status }}</p> -->
    <h2>Uploaded File List</h2>
    <FileTable :tables="tables" @update-tables="updateTables"/>
    <FooterComponent />
  </div>
</template>

<script>
import HeaderComponent from "./components/HeaderComponent.vue";
import UploadForm from "./components/UploadForm.vue";
import FileTable from "./components/FileTable.vue";
import FooterComponent from "./components/FooterComponent.vue";
import axios from "axios";
import { config } from "../config";

export default {
  name: "PaperManagement",
  components: {
    HeaderComponent,
    UploadForm,
    FileTable,
    FooterComponent,
  },
  data() {
    return {
      // file_name: null, //Goから受け取るアップロードされたファイル名
      // upload_status: null, //Goから受け取るアップロードステータス
      tables: [], // Goから受け取るテーブル一覧を格納するデータプロパティ
    };
  },
  created() {
    // MySQLに接続してテーブル一覧を取得する
    this.getDB();
  },
  methods: {
    updateTables() {
      // 子コンポーネントから受け取ったデータを、tablesプロパティにセット
      this.getDB();
    },
    updateUploadStatus() {
      // this.file_name = responseData.filename;
      // this.upload_status = responseData.status;
      this.getDB();
    },
    getDB: function(){
      axios.get(`${config.URL}:${config.PORT}/api/tables`).then((res) => {
        this.tables = res.data;
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
}
/* filename: 及び status: を装飾する */
p {
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: bold;
}
</style>
