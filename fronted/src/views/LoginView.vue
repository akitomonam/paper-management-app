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
        <router-link to="/signup">サインアップ</router-link>
    </div>
</template>

<script>
import axios from 'axios';
import { config } from '../../config';

export default {
    name: "LoginView",
    data() {
        return {
            username: "",
            password: "",
        }
    },
    methods: {
        async login() {
            // サーバーにログイン要求を送信する
            try {
                const response = await axios.post(`${config.URL}:${config.PORT}/api/login`, {
                    username: this.username,
                    password: this.password,
                }, {
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded',
                    },
                    // proxy: false //ローカルホストなのでプロキシを経由しない
                });
                // ログインが成功した場合
                if (response.data.success) {
                    alert(response.data.message);
                    // マイページ画面に遷移する
                    this.$router.push({ path: '/mypage' });
                } else {
                    // ログインが失敗した場合、エラーメッセージを表示する
                    alert(response.data.message);
                }
            } catch (error) {
                console.error(error);
            }
        },
    },
}
</script>
