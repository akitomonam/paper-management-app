<template>
    <div>
        <h1 style="text-align: center;">Registered Users List</h1>
        <div class="demo-type">
            <div v-for="table in tables" :key="table">
                <el-avatar > {{ table.Username }} </el-avatar>
            </div>
        </div>
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

<style scoped>
.demo-type {
    display: flex;
}

.demo-type>div {
    flex: 1;
    text-align: center;
}

.demo-type>div:not(:last-child) {
    border-right: 1px solid var(--el-border-color);
}
</style>
