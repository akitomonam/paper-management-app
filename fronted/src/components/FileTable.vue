<template>
    <table>
        <div class="draggable-container">
            <draggable v-model="localTables" group="people" item-key="ID" handle=".handle">
                <template #item="{ element }">
                    <tr style="border: solid 1px #000">
                        <div class="drag-item">
                            <img src="../assets/drag_drop_button.png" class="handle" />
                            <!-- {{ element.file_name }} -->
                            <!-- <router-link :to="'/papers/' + element.ID">{{ element.file_name }}</router-link> -->
                            <router-link :to="'/papers/' + element.ID">
                                {{ element.title ? element.title : element.file_name }}
                            </router-link>
                        </div>
                    </tr>
                </template>
            </draggable>
        </div>
    </table>
</template>

<script>
import draggable from "vuedraggable";
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
