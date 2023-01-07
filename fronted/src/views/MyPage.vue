<template>
    <div>
        <div>
            <UserProfileInfo :user="user" />
        </div>
        <div>
            <UploadForm @update-upload-status="updateTables" />
        </div>
        <UserPaperTableInfo ref="tables"/>
        <div>
            <h2>Edit Settings</h2>
            <form>
                <label>
                    Profile Image:
                    <input type="file" @change="onProfileFileChange" />
                </label>
                <br />
                <label>
                    username:
                    <input type="text" v-model="user.name" />
                </label>
                <br />
                <label>
                    self-introductory statement:
                    <textarea v-model="user.description"></textarea>
                </label>
                <br />
                <button type="submit">Save</button>
            </form>
        </div>
        <div>
            <h2>Account Management</h2>
            <button @click="deleteAccount">Delete account</button>
            <button v-if="user.isLocked" @click="unlockAccount">Unlock account</button>
            <button v-else @click="lockAccount">Lock account</button>
        </div>
    </div>
</template>

<script>
import axios from "axios";
import { config } from "../../config";
import UserProfileInfo from '../components/UserProfileInfo.vue'
import UploadForm from "../components/UploadForm.vue";
import UserPaperTableInfo from "../components/UserPaperTableInfo.vue";

export default {
    name: "MyPage",
    components: {
        UserProfileInfo,
        UploadForm,
        UserPaperTableInfo,
    },
    data() {
        return {
            user: {
                avatarUrl: '',
                name: '',
                description: '',
                isLocked: false,
            },
        }
    },
    created() {
        this.getUsersDB()
    },
    methods: {
        deleteAccount() {
            // アカウントを削除する処理を記述する
        },
        lockAccount() {
            this.user.isLocked = true;
            // アカウントをロックする処理を記述する
        },
        unlockAccount() {
            this.user.isLocked = false;
            // アカウントをアンロックする処理を記述する
        },
        onFileChange(event) {
            const file = event.target.files[0];
            if (!file) {
                return;
            }

            const reader = new FileReader();
            reader.onload = (e) => {
                this.user.avatarUrl = e.target.result;
            };
            reader.readAsDataURL(file);
        },
        getUsersDB: function () {
            const sessionToken = localStorage.getItem('sessionToken');
            console.log("sessionToken:", sessionToken)
            axios.get(`${config.URL}:${config.PORT}/api/userinfo?sessionToken=${sessionToken}`).then((res) => {
                console.log("res.data:", res.data)
                this.user.name = res.data.Username;
                console.log("user.name:", this.user.name)
            });
        },
        updateTables(favorite = false) {
            this.$refs.tables.updateTables(favorite)
        },
    },
}
</script>

<style>
h1 {
    text-align: center;
}

div {
    margin: 20px;
}

img {
    width: 100px;
    height: 100px;
    border-radius: 50%;
}

ul {
    list-style: none;
    padding: 0;
}

li {
    margin: 10px 0;
}
</style>
