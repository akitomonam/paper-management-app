<template>
    <el-card class="box-card login">
        <div class="clearfix">
            <span>Login</span>
            <el-button style="float: right; padding: 3px 0" type="text" @click="register">
                <el-icon el-icon--left>
                    <UserFilled />
                </el-icon>
                Register
            </el-button>
        </div>
        <el-form :model="form" label-width="120px" style="max-width: 460px">
            <el-form-item label="User name">
                <el-input v-model="username" />
            </el-form-item>
            <el-form-item label="Password">
                <el-input type="password" autocomplete="off" v-model="password" />
            </el-form-item>
            <el-form-item>
                <div style="margin: 0 auto;">
                    <el-button type="primary" @click="login">Login</el-button>

                </div>
                <!-- <el-button @click="resetForm(ruleFormRef)">Reset</el-button> -->
            </el-form-item>
            <!-- <router-link to="/signup">Register</router-link> -->
        </el-form>
    </el-card>
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
        register() {
            this.$router.push({ path: '/signup' });
        }
    }
}
</script>

<style scoped>
.box-card {
    width: 480px;
}

.login {
    position: relative;
    top: 100px;
    right: 0px;
    bottom: 0px;
    left: 0px;
    margin: auto;
}
</style>
