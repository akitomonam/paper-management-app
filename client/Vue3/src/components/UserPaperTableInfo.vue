<template>
<div>
    <div style="text-align: center; margin-top: 20px;">
        <el-card class="box-card">
            <template #header>
                <div class="card-header">
                    <span>My Uploaded Paper List</span>
                    <!-- <el-button class="button" text>Operation button</el-button> -->
                </div>
            </template>
            <FileTable :tables="tables" @update-tables="updateTables" />
        </el-card>
        <el-card class="box-card">
            <template #header>
                <div class="card-header">
                    <span>My Favorite Paper List</span>
                    <!-- <el-button class="button" text>Operation button</el-button> -->
                </div>
            </template>
            <FileTable :tables="favorite_tables" @update-tables="updateTables(true)" />
        </el-card>
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
