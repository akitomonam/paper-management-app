<template>
    <div>
        <h1>マイページ</h1>
        <div>
            <h2>プロフィール情報</h2>
            <div>
                <img src="{{ user.avatarUrl }}" alt="アバター画像" />
                <p>{{ user.name }}</p>
                <p>{{ user.description }}</p>
            </div>
        </div>
        <div>
            <h2>アクティビティの履歴</h2>
            <ul>
                <li v-for="activity in activities" :key="activity.id">{{ activity.content }}</li>
            </ul>
        </div>
        <div>
            <h2>設定の編集</h2>
            <form>
                <label>
                    プロフィール画像:
                    <input type="file" @change="onProfileFileChange" />
                </label>
                <br />
                <label>
                    ユーザー名:
                    <input type="text" v-model="user.name" />
                </label>
                <br />
                <label>
                    自己紹介文:
                    <textarea v-model="user.description"></textarea>
                </label>
                <br />
                <button type="submit">保存</button>
            </form>
        </div>
        <div>
            <h2>アカウントの管理</h2>
            <button @click="deleteAccount">アカウントを削除する</button>
            <button v-if="user.isLocked" @click="unlockAccount">アカウントをアンロックする</button>
            <button v-else @click="lockAccount">アカウントをロックする</button>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            user: {
                avatarUrl: '',
                name: '',
                description: '',
                isLocked: false,
            },
            activities: [],
        }
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
