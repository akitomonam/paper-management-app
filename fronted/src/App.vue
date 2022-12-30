<template>
  <div>
    <button class="settings-button" @click="changeOpen">
      &#9776;
    </button>
    <HeaderComponent />
    <SideBar v-bind:isOpen="isOpen" @changeOpen="changeOpen" />
    <UploadForm @update-upload-status="updateUploadStatus" />
    <!-- <p>filename:{{ file_name }}</p>
    <p>status:{{ upload_status }}</p> -->
    <h2>Uploaded File List</h2>
    <FileTable :tables="tables" @update-tables="updateTables" />
    <FooterComponent />
  </div>
</template>

<script>
import HeaderComponent from "./components/HeaderComponent.vue";
import SideBar from "./components/SideBar.vue";
import UploadForm from "./components/UploadForm.vue";
import FileTable from "./components/FileTable.vue";
import FooterComponent from "./components/FooterComponent.vue";
import axios from "axios";
import { config } from "../config";

export default {
  name: "PaperManagement",
  components: {
    HeaderComponent,
    SideBar,
    UploadForm,
    FileTable,
    FooterComponent,
  },
  data() {
    return {
      // file_name: null, //Goから受け取るアップロードされたファイル名
      // upload_status: null, //Goから受け取るアップロードステータス
      isOpen: false, // サイドバーの表示状態を管理するデータプロパティ
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
    getDB: function () {
      axios.get(`${config.URL}:${config.PORT}/api/tables`).then((res) => {
        this.tables = res.data;
      });
    },
    changeOpen() {
      // サイドバーの表示状態を反転する
      this.isOpen = !this.isOpen;
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

.settings-button {
  position: absolute;
  /* 設定ボタンを絶対配置にする */
  top: 0px;
  /* 設定ボタンをページの上部に配置する */
  right: 10px;
  /* 設定ボタンをページの右上に配置する */
  background-color: #fff;
  /* 設定ボタンの背景色を指定 */
  border: none;
  /* 設定ボタンの枠線を非表示にする */
  cursor: pointer;
  /* マウスカーソルをポインターにする */
  font-size: 20px;
  /* 設定ボタンの文字サイズを指定 */
  z-index: 2;
  /* 設定ボタンを前面に表示する */
}

.settings-button:hover {
  /* 設定ボタンにマウスが乗った時の装飾を追加 */
  background-color: #eee;
  /* 設定ボタンの背景色を指定 */
}
</style>
