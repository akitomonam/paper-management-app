<template>
<div>
    <div>
        <h2>My Uploaded File List</h2>
        <FileTable :tables="tables" @update-tables="updateTables" />
        <h2>My Favorite File List</h2>
        <FileTable :tables="favorite_tables" @update-tables="updateTables(true)" />
    </div>
</div>
</template>

<script>
import axios from "axios";
import { config } from "../../config";
import FileTable from "../components/FileTable.vue";
export default {
    name: "UserPaperTableInfo",
    components: {
        FileTable,
    },
    data() {
        return {
            tables: [],
            favorite_tables: []
        }
    },
    created() {
        this.updateTables()
        this.updateTables(true)
    },
    methods: {
        updateTables(favorite = false) {
            this.getDB(favorite);
        },
        getDB: function (favorite) {
            const sessionToken = localStorage.getItem('sessionToken');
            axios.get(`${config.URL}:${config.PORT}/api/tables?sessionToken=${sessionToken}&favorite=${favorite}`).then((res) => {
                if (favorite) {
                    this.favorite_tables = res.data
                } else {
                    this.tables = res.data;
                }
            });
        },
    }
}
</script>

<style>
</style>
