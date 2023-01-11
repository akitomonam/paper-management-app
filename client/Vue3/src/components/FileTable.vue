<template>
    <div>
        <!-- <table>
            <div class="draggable-container">
                <draggable v-model="localTables" group="people" item-key="ID" handle=".handle">
                    <template #item="{ element }">
                        <tr style="border: solid 1px #000">
                            <div class="drag-item">
                                <img src="../assets/drag_drop_button.png" class="handle" />
                                <router-link :to="'/papers/' + element.ID">
                                    {{ element.title ? element.title : element.file_name }}
                                </router-link>
                            </div>
                        </tr>
                    </template>
                </draggable>
            </div>
        </table> -->
        <vue-good-table :columns="columns" :rows="localTables" :search-options="{
            enabled: true,
            skipDiacritics: true,
            // searchFn: mySearchFn,
            placeholder: 'Search this table',
            // externalQuery: searchQuery
        }" v-on:row-click="onRowClick"/>
    </div>
</template>

<script>
// import draggable from "vuedraggable";
import 'vue-good-table-next/dist/vue-good-table-next.css'
import { VueGoodTable } from 'vue-good-table-next';
export default {
    name: "FileTable",
    components: {
        // draggable,
        VueGoodTable
    },
    props: {
        tables: {
            type: Array,
            required: true,
        },
    },
    data() {
        return {
            localTables: this.tables,
            columns: [
                {
                    label: 'ID',
                    field: 'ID',
                },
                {
                    label: 'Title',
                    field: 'title',
                },
                {
                    label: 'author',
                    field: 'author',
                },
                {
                    label: 'publisher',
                    field: 'publisher',
                },
                {
                    label: 'year',
                    field: 'year',
                },
                {
                    label: 'Filename',
                    field: 'file_name',
                },
            ],
        };
    },
    created() {
        this.localTables = this.tables;
    },
    methods: {
        onRowClick(params) {
            this.$router.push({ path: '/papers/' + params.row.ID });
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

<style scoped>
/* アップロード済みファイル一覧を装飾する */

/* table {
    border-collapse: collapse;
    margin: 0 auto;
}

tr:nth-child(even) {
    background-color: #f2f2f2;
}

tr:hover {
    background-color: #ddd;
} */

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
