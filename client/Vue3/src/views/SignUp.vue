<template>
    <div class="Registration">
        <div class="Registration-triangle"></div>
        <h2 class="Registration-header">Registration</h2>
        <form class="Registration-container">
            <p><input type="username" placeholder="Username" v-model="username"></p>
            <p><input type="password" placeholder="Password" v-model="password"></p>
            <p><input type="submit" value="Registration" @click.prevent="signup"></p>
        </form>
        <p><router-link to="/login">login</router-link></p>
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

<style scoped>
body {
    background: rgba(255, 255, 255, 0);
    font-family: 'Open Sans', sans-serif;
}

.Registration {
    width: 400px;
    margin: 16px auto;
    font-size: 16px;
}

/* Reset top and bottom margins from certain elements */
.Registration-header,
.Registration p {
    margin-top: 0;
    margin-bottom: 0;
}

/* The triangle form is achieved by a CSS hack */
.Registration-triangle {
    width: 0;
    margin-right: auto;
    margin-left: auto;
    border: 12px solid transparent;
    border-bottom-color: #28d;
}

.Registration-header {
    background: #28d;
    padding: 20px;
    font-size: 1.4em;
    font-weight: normal;
    text-align: center;
    text-transform: uppercase;
    color: #fff;
}

.Registration-container {
    background: #ebebeb;
    /* padding: 20px; */
}

/* Every row inside .Registration-container is defined with p tags */
.Registration p {
    padding: 12px;
}

.Registration input {
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

.Registration input[type="email"],
.Registration input[type="password"] {
    background: #fff;
    border-color: #bbb;
    color: #555;
}

/* Text fields' focus effect */
.Registration input[type="email"]:focus,
.Registration input[type="password"]:focus {
    border-color: #888;
}

.Registration input[type="submit"] {
    background: #28d;
    border-color: transparent;
    color: #fff;
    cursor: pointer;
}

.Registration input[type="submit"]:hover {
    background: #17c;
}

/* Buttons' focus effect */
.Registration input[type="submit"]:focus {
    border-color: #05a;
}
</style>
