<template>
    <div style="text-align: center;">
        <h1>{{ user.name }}'s Page</h1>
        <el-avatar> {{ user.name }} </el-avatar>
        <!-- <h2>Profile Info</h2>
        <div>
            <img :src="user.avatarUrl" alt="Profile Image" />
            <p>username: {{ user.name }}</p>
            <p>{{ user.description }}</p>
        </div> -->
    </div>
</template>

<script>
export default {
    name: "UserProfileInfo",
    components: {
    },
    data() {
        return {
        }
    },
    props: {
        user: {
            type: Object,
            required: true,
        },
    }
}
</script>

<style scoped>
div {
    margin: 20px;
}
</style>
