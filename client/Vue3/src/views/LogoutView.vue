<template>
    <el-card class="box-card logout">
        <template #header>
            <div class="card-header">
                <span>Logout</span>
            </div>
        </template>
        <el-form :model="form" label-width="120px" style="max-width: 460px; text-align: center;">
            <el-form-item>
                <div style="margin: auto;">
                    <el-button type="primary" @click="logout">Logout</el-button>
                </div>
            </el-form-item>
        </el-form>
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
