<template>
    <div>
        <el-card class="box-card">
            <template #header>
                <div class="card-header">
                    <span>Comment</span>
                    <!-- <el-button class="button" text>Operation button</el-button> -->
                </div>
            </template>
            <el-form>
                <el-form-item>
                    <el-input v-model="message" placeholder="Enter your message" clearable
                        @keydown.enter.prevent="postMessage"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button @click="postMessage" type="primary">Post</el-button>
                </el-form-item>
            </el-form>
            <el-card v-for="message in reverseItems" :key="message.CreatedAt" class="bulletin-card">
                <div style="display:flex;">
                    <el-avatar> {{ message.Username }} </el-avatar>
                    <div style="margin-left: 10px;">{{ message.Content }}</div>
                </div>
                <p style="font-size: 12px;">{{ message.CreatedAt }}</p>
            </el-card>
        </el-card>
    </div>
</template>

<script>
import axios from 'axios'
import { config } from "../../config";

export default {
    props: {
        "paper_id": {
            required: true
        }
    },
    data() {
        return {
            messages: [],
            message: '',
        }
    },
    computed: {
        reverseItems() {
            return this.messages.slice().reverse();
        },
    },
    methods: {
        async fetchMessages() {
            axios.get(`${config.URL}:${config.PORT}/api/comment_preview`,
                {
                    params: {
                        paperId: this.paper_id,
                    }
                })
                .then(response => {
                    console.log("messages", response.data)
                    this.messages = response.data
                    return this.messages
                })
                .catch((error) => {
                    console.log("コメント一覧取得APIでエラーが発生しました");
                    console.error(error);
                    alert("コメント一覧取得APIでエラーが発生しました")
                });
        },
        async postMessage() {
            console.log("paper_id", this.paper_id)
            axios.post(`${config.URL}:${config.PORT}/api/comment_add`,
                {
                    sessionToken: localStorage.getItem('sessionToken'),
                    paperId: this.paper_id,
                    comments: this.message

                },
                {
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded',
                    },
                })
                .then(response => {
                    console.log(response.data)
                    this.fetchMessages()
                })
                .catch((error) => {
                    console.log("コメント投稿APIでエラーが発生しました");
                    console.error(error);
                    alert("コメント投稿APIでエラーが発生しました")
                });
            this.message = ""
        },
    },
    watch: {
        paper_id() {
            this.fetchMessages()
        }
    },
}
</script>

<style>
.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.bulletin-card {
    /* width: auto; */
    margin: 1px 15px;
    /* position: relative; */
}

.box-card {
    width: auto;
    margin: 10px;
    position: relative;
}
</style>
