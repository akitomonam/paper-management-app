<template>
    <div>
        <h1>My Page</h1>
        <div>
            <h2>Profile Info</h2>
            <div>
                <img src="{{ user.avatarUrl }}" alt="Profile Image" />
                <p>username:{{ user.name }}</p>
                <p>{{ user.description }}</p>
            </div>
        </div>
        <div>
            <h2>My Uploaded File List</h2>
            <FileTable :tables="tables" @update-tables="updateTables" />
            <h2>My Favorite File List</h2>
            <FileTable :tables="favorite_tables" @update-tables="updateTables(true)" />
        </div>
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
import FileTable from "../components/FileTable.vue";

export default {
    name: "MyPage",
    components: {
        FileTable,
    },
    data() {
        return {
            user: {
                avatarUrl: '',
                name: '',
                description: '',
                isLocked: false,
            },
            tables: [],
            favorite_tables: []
        }
    },
    created() {
        this.getUsersDB()
        this.updateTables()
        this.updateTables(true)
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
        updateTables(favorite=false) {
            // 子コンポーネントから受け取ったデータを、tablesプロパティにセット
            this.getDB(favorite);
        },
        getDB: function (favorite) {
            const sessionToken = localStorage.getItem('sessionToken');
            axios.get(`${config.URL}:${config.PORT}/api/tables?sessionToken=${sessionToken}&favorite=${favorite}`).then((res) => {
                if (favorite) {
                    this.favorite_tables = res.data
                } else {
                    this.tables = res.data;
                }
            });
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
