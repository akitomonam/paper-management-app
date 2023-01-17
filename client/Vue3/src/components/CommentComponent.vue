Copy code
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
                    <el-input v-model="message" placeholder="Enter your message"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button @click="postMessage" type="primary">Post</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card v-for="message in messages" :key="message.id" class="bulletin-card">
            <p>{{ message.text }}</p>
            <p>{{ message.date }}</p>
            <el-card v-for="comment in message.comments" :key="comment.id" class="comment-card">
                <p>{{ comment.text }}</p>
                <p>{{ comment.date }}</p>
            </el-card>
        </el-card>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    data() {
        return {
            messages: [],
            message: '',
        }
    },
    created() {
        this.fetchMessages()
    },
    methods: {
        async postMessage() {
            // send the message to the backend
            let response = await axios.post('/api/messages', { text: this.message })
            if (response.data) {
                this.message = ''
                this.fetchMessages()
            }
        },
        async fetchMessages() {
            // fetch messages and their comments from the backend
            let response = await axios.get('/api/messages')
            if (response.data) {
                this.messages = response.data
            }
        },
    },
}
</script>

<style>
.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.box-card {
    width: auto;
    margin: 10px;
    position: relative;
}
</style>
