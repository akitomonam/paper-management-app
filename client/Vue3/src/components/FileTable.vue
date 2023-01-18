<template>
    <div>
        <el-table :data="filterTableData" style="width: 100%">
            <el-table-column label="ID" prop="ID" sortable/>
            <el-table-column label="Title" prop="title" sortable/>
            <el-table-column label="Author" prop="author" sortable/>
            <el-table-column label="Publisher" prop="publisher" sortable/>
            <el-table-column label="Year" prop="year" sortable/>
            <el-table-column label="File name" prop="file_name" sortable/>
            <el-table-column align="right">
                <template #header>
                    <el-input v-model="search" size="small" placeholder="Type to search"/>
                </template>
                <template #default="scope">
                    <el-button size="small" @click="handleEdit(scope.$index, scope.row)"><el-icon>
                        <Plus />
                    </el-icon>More info</el-button>
                    <el-button size="small" type="info" @click="showFile(scope.$index, scope.row)"><el-icon el-icon--left>
                        <View />
                    </el-icon>Preview</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
// import draggable from "vuedraggable";
import axios from 'axios'
// import 'vue-good-table-next/dist/vue-good-table-next.css'
import { config } from "../../config";
export default {
    name: "FileTable",
    components: {
        // draggable,
    },
    props: {
        tables: {
            type: Array,
            required: true,
        },
    },
    data() {
        return {
            search: '',
            localTables: this.tables,
        };
    },
    created() {
        this.localTables = this.tables;
    },
    computed: {
        filterTableData() {
            return this.tables.filter(
                (data) =>
                    !this.search ||
                    data.title.toLowerCase().includes(this.search.toLowerCase()) ||
                    data.author.toLowerCase().includes(this.search.toLowerCase()) ||
                    data.publisher.toLowerCase().includes(this.search.toLowerCase()) ||
                    data.file_name.toLowerCase().includes(this.search.toLowerCase()) ||
                    data.year.toString().toLowerCase().includes(this.search.toString().toLowerCase())
            );
        },
    },
    methods: {
        onRowClick(params) {
            this.$router.push({ path: '/papers/' + params.row.ID });
        },
        handleEdit(index, row) {
            // code to handle edit
            console.log(index, row)
            this.$router.push({ path: '/papers/' + row.ID });
        },
        handleDelete(index, row) {
            // code to handle delete
            console.log(index, row)
            console.log("row ID:", row.ID)
        },
        // tableをクリックした際に実行される処理
        async showFile(index, row) {
            console.log("click table!!!");
            console.log("targetFileId", row.ID); // 目的の文字列を出力する
            // サーバーに保管されているファイルをプレビューする
            axios
                .get(
                    `${config.URL}:${config.PORT}/api/preview?fileId=${row.ID}`
                )
                .then((response) => {
                    // プレビューするファイルのURLを取得する
                    let filepath = response.data.fileUrl;
                    // プレビューするファイルのURLをもとに、新しいタブを開く
                    filepath = filepath.replace(/^\./, "");
                    window.open(`${config.URL}:${config.PORT}${filepath}`, "_blank");
                })
                .catch((error) => {
                    console.log("ファイルプレビューAPIでエラーが発生しました");
                    console.error(error);
                });
        },
    },
    watch: {
        // tablesプロパティが更新されたら、localTablesも更新
        tables(newTables) {
            this.localTables = newTables;
        },
    },
};
</script>

<style scoped>
.handle {
    /* background: #3d3c3c17; */
    border-radius: 50%;
    width: 20px;
    height: 20px;
    color: #fff;
    /* font-size: 10px; */
    text-align: center;
    /* line-height: 110px; */
}
</style>
