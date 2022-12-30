<template>
    <div>
        <UploadForm @update-upload-status="updateUploadStatus" />
        <h2>Uploaded File List</h2>
        <FileTable :tables="tables" @update-tables="updateTables" />
    </div>
</template>

<script>
import UploadForm from "../components/UploadForm.vue";
import FileTable from "../components/FileTable.vue";
import axios from "axios";
import { config } from "../../config";

export default {
    name: "PaperManagement",
    components: {
        UploadForm,
        FileTable,
    },
    data() {
        return {
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
            this.getDB();
        },
        getDB: function () {
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
</style>
