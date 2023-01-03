<template>
    <div>
        <h1>ユーザー登録</h1>
        <form>
            <label for="username">ユーザー名</label>
            <input type="text" v-model="username" id="username" />
            <br />
            <label for="password">パスワード</label>
        <input type="password" v-model="password" id="password" />
        <br />
        <button @click.prevent="signup">サインアップ</button>
    </form>
</div>
</template>
<script>
export default {
    data() {
        return {
            username: '',
            password: '',
        };
    },
    methods: {
        signup() {
            // ここでユーザー名とパスワードを使用してサインアップ処理を行う
        },
    },
};
</script>
<style>

</style>
