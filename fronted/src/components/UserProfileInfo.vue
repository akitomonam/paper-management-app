<template>
    <div>
        <h1>{{ user.name }}'s Page</h1>
        <h2>Profile Info</h2>
        <div>
            <img :src="user.avatarUrl" alt="Profile Image" />
            <p>username: {{ user.name }}</p>
            <p>{{ user.description }}</p>
        </div>
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
</style>
