<template>
    <div>
        <UserProfileInfo :user="user" />
        <UploadForm @update-upload-status="updateTables" />
        <UserPaperTableInfo ref="tables" />
        <UserPrifileInfoEdit :user="user" />
    </div>
</template>

<script>
import axios from "axios";
import { config } from "../../config";
import UserProfileInfo from '../components/UserProfileInfo.vue'
import UploadForm from "../components/UploadForm.vue";
import UserPaperTableInfo from "../components/UserPaperTableInfo.vue";
import UserPrifileInfoEdit from "../components/UserPrifileInfoEdit.vue";

export default {
    name: "MyPage",
    components: {
        UserProfileInfo,
        UploadForm,
        UserPaperTableInfo,
        UserPrifileInfoEdit
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
