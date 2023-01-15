<template>
    <el-card class="box-card register">
        <template #header>
            <div class="card-header">
                <span>Registration</span>
                <el-button type="text" @click="login">
                    <el-icon el-icon--left>
                        <UserFilled />
                    </el-icon>
                    Login
                </el-button>
            </div>
        </template>
        <el-form :model="form" label-width="120px" style="max-width: 460px">
            <el-form-item label="User name">
                <el-input v-model="username" />
            </el-form-item>
            <el-form-item label="Password">
                <el-input type="password" autocomplete="off" v-model="password" />
            </el-form-item>
            <el-form-item>
                <div style="margin: 0 auto;">
                    <el-button type="primary" @click="signup">Register</el-button>

                </div>
            </el-form-item>
        </el-form>
    </el-card>
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
        login() {
            this.$router.push({ path: '/login' });
        }
    },
};
</script>

<style scoped>
.box-card {
    width: 480px;
}
.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
.register {
    position: relative;
    top: 100px;
    right: 0px;
    bottom: 0px;
    left: 0px;
    margin: auto;
}
</style>
