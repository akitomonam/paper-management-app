<template>
    <div>
        <h1>User Registration Page</h1>
        <form>
            <label for="username">username</label>
            <input type="text" v-model="username" id="username" />
            <br />
            <label for="password">password</label>
        <input type="password" v-model="password" id="password" />
        <br />
        <button @click.prevent="signup">signup</button>
    </form>
</div>
</template>
<script>
import axios from 'axios';
import { config } from '../../config';

export default {
    data() {
        return {
            username: '',
            password: '',
        };
    },
    methods: {
        async signup() {
            try {
                const response = await axios.post(`${config.URL}:${config.PORT}/api/signup`, {
                    username: this.username,
                    password: this.password,
                }, {
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded',
                    },
                });
                // サインアップが成功した場合
                if (response.data.success) {
                    alert(response.data.message);
                    // ログイン画面に遷移する
                    this.$router.push({ path: '/login' });
                } else {
                    // ログインが失敗した場合、エラーメッセージを表示する
                    alert(response.data.message);
                }
            } catch (error) {
                console.error(error);
            }
        },
    },
};
</script>

<style>

</style>
