<template>
    <table>
        <div class="draggable-container">
            <draggable v-model="localTables" group="people" item-key="ID" handle=".handle">
                <template #item="{ element }">
                    <tr style="border: solid 1px #000">
                        <div class="drag-item">
                            <img src="../assets/drag_drop_button.png" class="handle" />
                            {{ element.filename }}
                            <button @click="showFile(element.filepath)">Preview</button>
                            <button @click="deleteFile(element.ID)">Delete</button>
                        </div>
                    </tr>
                </template>
            </draggable>
        </div>
    </table>
</template>

<script>
import axios from "axios";
import draggable from "vuedraggable";
import { config } from "../../config";
export default {
    name: "FileTable",
    components: {
        draggable,
    },
    props: {
        tables: {
            type: Array,
            required: true,
        },
    },
    data() {
        return {
            // プロパティを変更するためのデータプロパティを用意
            // 初期値は、propsから受け取った値を設定
            localTables: this.tables,
        };
    },
    created() {
        this.localTables = this.tables;
    },
    methods: {
        // tableをクリックした際に実行される処理
        async showFile(targetFilepath) {
            console.log("click table!!!");
            console.log("targetFilepath", targetFilepath); // 目的の文字列を出力する
            // サーバーに保管されているファイルをプレビューする
            axios
                .get(
                    `${config.URL}:${config.PORT}/api/preview?fileName=${targetFilepath}`
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
        async deleteFile(targetFileId) {
            console.log("targetFileId:", targetFileId)
            try {
                const response = await axios.get(`${config.URL}:${config.PORT}/api/delete?fileId=${targetFileId}`)
                if (response.data) { // レスポンスボディが存在する場合
                    console.log("response.data", response.data);
                    if (response.data.result == "true") {
                        alert("ファイル削除に成功しました。")
                        this.$emit('update-tables') // 親コンポーネントに発火
                    } else {
                        alert("ファイル削除に失敗しました。")
                    }
                } else {
                    console.error('レスポンスボディが存在しません')
                }
            } catch {
                console.log('File delete Failed');
            }
        }
    },
    watch: {
        // tablesプロパティが更新されたら、localTablesも更新
        tables(newTables) {
            this.localTables = newTables;
        },
    },
};
</script>

<style>
/* アップロード済みファイル一覧を装飾する */
table {
    border-collapse: collapse;
    margin: 0 auto;
}

tr:nth-child(even) {
    background-color: #f2f2f2;
}

tr:hover {
    background-color: #ddd;
}

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
