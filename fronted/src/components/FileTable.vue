<template>
  <table>
    <div
      class="draggable-container"
      style="display: flex; justify-content: center"
    >
      <draggable v-model="localTables" group="people" item-key="ID" handle=".handle">
        <template #item="{ element }">
          <tr style="border: solid 1px #000">
            <div class="drag-item">
              <span class="handle">・</span>
              {{ element.filename }}
              <button @click="showFile(element.filepath)">Preview</button>
            </div>
          </tr>
        </template>
      </draggable>
    </div>
  </table>
</template>
  
<script>
import axios from 'axios';
import draggable from "vuedraggable";
import { config } from '../../config';
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
    // 親コンポーネントにイベントを発火させる
    updateTables() {
        this.$emit("update-tables", this.localTables);
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
  watch: {
    // tablesプロパティが更新されたら、localTablesも更新
    tables(newTables) {
      this.localTables = newTables;
    },
    // localTablesが更新されたら、updateTablesを実行
    localTables() {
      this.updateTables();
    },
  },
};
</script>

<style>
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
</style>