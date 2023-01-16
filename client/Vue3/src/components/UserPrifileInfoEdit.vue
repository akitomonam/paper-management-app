<template>
    <div style="text-align: center;">
        <!-- <div>Account Management</div> -->
        <el-button size="small" type="danger" @click="deleteAccount"><el-icon el-icon--left>
                <Delete />
            </el-icon>Delete Account</el-button>
    </div>
</template>

<script>
export default {
    name: "UserPrifileInfoEdit",
    props: {
        user: {
            type: Object,
            required: true,
        },
    },
    methods: {
        deleteAccount() {
            // アカウントを削除する処理を記述する
        },
    }
}
</script>

<style scoped>

</style>
