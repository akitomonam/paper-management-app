<template>
    <!-- tablesリスト中の辞書に含まれるnameを一覧表示する -->
    <div>
        <h1>Registered Users List</h1>
        <li v-for="table in tables" :key="table">
            {{ table.Username }}
        </li>
    </div>
</template>

<script>
import axios from "axios";
import { config } from "../../config";
export default {
    name: "FileTable",
    data() {
        return {
            tables :[]
        };
    },
    created() {
        this.getUserList()
    },
    methods: {
        getUserList: function () {
            axios.get(`${config.URL}:${config.PORT}/api/userlist`).then((res) => {
                console.log("res.data:", res.data)
                this.tables = res.data;
            });
        },
    }
}
</script>

<style>
</style>
