<template>
    <div>
        <h1>Logout</h1>
        <!-- ログアウトボタンを表示する -->
        <button @click="logout">ログアウト</button>
    </div>
</template>
<script>
export default {
    name: "MyPage",
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

<style>

</style>
