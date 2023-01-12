<template>
<div class="login">
    <div class="login-triangle"></div>

    <h2 class="login-header">Log in</h2>

    <form class="login-container">
        <p><input type="username" placeholder="Username" v-model="username"></p>
        <p><input type="password" placeholder="Password" v-model="password" ></p>
        <p><input type="submit" value="Log in" @click.prevent="login"></p>
    </form>
    <p><router-link to="/signup">Register</router-link></p>
</div>
</template>

<script>
import axios from 'axios';
import { config } from '../../config';

export default {
    name: 'LoginView',
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
                    // If the login is successful, store the session token in local storage
                    console.log("response.data.token:", response.data.token)
                    localStorage.setItem('sessionToken', response.data.token);
                    // alert(response.data.message);
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
    }
}
</script>

<style scoped>
@import url(https://fonts.googleapis.com/css?family=Open+Sans:400,700);

body {
    background: rgba(255, 255, 255, 0);
    font-family: 'Open Sans', sans-serif;
}

.login {
    width: 400px;
    margin: 16px auto;
    font-size: 16px;
}

/* Reset top and bottom margins from certain elements */
.login-header,
.login p {
    margin-top: 0;
    margin-bottom: 0;
}

/* The triangle form is achieved by a CSS hack */
.login-triangle {
    width: 0;
    margin-right: auto;
    margin-left: auto;
    border: 12px solid transparent;
    border-bottom-color: #28d;
}

.login-header {
    background: #28d;
    padding: 20px;
    font-size: 1.4em;
    font-weight: normal;
    text-align: center;
    text-transform: uppercase;
    color: #fff;
}

.login-container {
    background: #ebebeb;
    padding: 12px;
}

/* Every row inside .login-container is defined with p tags */
.login p {
    padding: 12px;
}

.login input {
    box-sizing: border-box;
    display: block;
    width: 100%;
    border-width: 1px;
    border-style: solid;
    padding: 16px;
    outline: 0;
    font-family: inherit;
    font-size: 0.95em;
}

.login input[type="email"],
.login input[type="password"] {
    background: #fff;
    border-color: #bbb;
    color: #555;
}

/* Text fields' focus effect */
.login input[type="email"]:focus,
.login input[type="password"]:focus {
    border-color: #888;
}

.login input[type="submit"] {
    background: #28d;
    border-color: transparent;
    color: #fff;
    cursor: pointer;
}

.login input[type="submit"]:hover {
    background: #17c;
}

/* Buttons' focus effect */
.login input[type="submit"]:focus {
    border-color: #05a;
}
</style>
