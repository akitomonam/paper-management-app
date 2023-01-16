<template>
    <div>
        <!-- <h1>Home</h1> -->
        <UploadForm @update-upload-status="updateUploadStatus" style="margin-top:20px"/>
        <el-card class="box-card">
            <template #header>
                <div class="card-header">
                    <span>Paper List</span>
                    <!-- <el-button class="button" text>Operation button</el-button> -->
                </div>
            </template>
            <FileTable :tables="tables" @update-tables="updateTables" />
        </el-card>
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

<style scoped>
.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.text {
    font-size: 14px;
}

.item {
    margin-bottom: 18px;
}

.box-card {
    width: auto;
    margin: 30px;
}
</style>
