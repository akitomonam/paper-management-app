<template>
    <el-card class="box-card logout">
        <template #header>
            <div class="card-header">
                <span>Logout</span>
            </div>
        </template>
                <div style="margin: auto; text-align: center;">
                    <el-button type="primary" @click="logout">Logout</el-button>
                </div>
    </el-card>
</template>
<script>
export default {
    name: "LogOut",
    methods: {
        logout() {
            // ローカルストレージからセッショントークンを削除する
            localStorage.removeItem('sessionToken');
            // ログイン画面にリダイレクトする
            this.$router.push({ path: '/login' });
        },
    },
}
</script>

<style scoped>
.box-card {
    width: 480px;
}

.logout {
    position: relative;
    top: 100px;
    right: 0px;
    bottom: 0px;
    left: 0px;
    margin: auto;
}
</style>
