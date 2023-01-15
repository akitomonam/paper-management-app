<template>
    <div>
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
                    <!-- <input type="text" v-model="user.name" /> -->
                    <input type="text"/>
                </label>
                <br />
                <label>
                    self-introductory statement:
                    <!-- <textarea v-model="user.description"></textarea> -->
                    <textarea ></textarea>
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
        lockAccount() {
            // this.user.isLocked = true;
            // アカウントをロックする処理を記述する
        },
        unlockAccount() {
            // this.user.isLocked = false;
            // アカウントをアンロックする処理を記述する
        },
        onFileChange(event) {
            const file = event.target.files[0];
            if (!file) {
                return;
            }

            const reader = new FileReader();
            // reader.onload = (e) => {
            reader.onload = () => {
                // this.user.avatarUrl = e.target.result;
            };
            reader.readAsDataURL(file);
        },
    }
}
</script>

<style scoped>

</style>
