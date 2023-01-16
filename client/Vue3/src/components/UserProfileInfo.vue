<template>
    <div style="text-align: center;">
        <el-header>
            <div style="font-size:xx-large">{{ user.name }}'s Page</div>
        </el-header>
            <el-avatar size="large"> {{ user.name }} </el-avatar>
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
