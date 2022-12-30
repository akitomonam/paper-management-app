<!-- ログイン画面を表示するコンポーネント -->
<template>
    <div>
        <form>
            <label>ユーザー名:</label>
            <input type="text" v-model="username" />
            <br />
            <label>パスワード:</label>
            <input type="password" v-model="password" />
            <br />
            <button type="submit" @click.prevent="login">ログイン</button>
        </form>
    </div>
</template>

<script>
export default {
    name: "LoginView",
    data() {
        return {
            username: "",
            password: "",
        }
    },
    methods: {
        login() {
            // ログイン処理を行う
            // 処理が成功したら、マイページ画面に遷移する
            this.$router.push({ path: '/mypage' });
        },
    },
}
</script>
