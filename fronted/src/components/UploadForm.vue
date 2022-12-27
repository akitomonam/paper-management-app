<template>
    <form @submit.prevent="uploadFile">
        <input type="file" ref="fileInput" />
        <button type="submit">Upload</button>
    </form>
</template>
  
<script>
import axios from 'axios';
import { config } from '../../config';

export default {
    name: 'UploadForm',
    components: {
    },
    data() {
        return {
        };
    },
    created() {
    },
    methods: {
        async uploadFile() {
            const file = this.$refs.fileInput.files[0];
            const formData = new FormData();
            formData.append('file', file);
            try {
                const response = await axios.post(`${config.URL}:${config.PORT}/upload/file`, formData, {
                    headers: {
                        'Content-Type': 'multipart/form-data'
                    },
                    proxy: false //ローカルホストなのでプロキシを経由しない
                })
                if (response.data) { // レスポンスボディが存在する場合
                    this.$emit('update-upload-status', response.data) // レスポンスボディのfilenameを参照して、responseMessageに代入する
                } else {
                    console.error('レスポンスボディが存在しません')
                }
            } catch (error) {
                console.log('File uploaded Failed');
                console.error(error);
            }
        },
    }
}
</script>

<style>
/* アップロードフォームを装飾する */
form {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
}

/* アップロードボタンを装飾する */
button[type="submit"] {
    background-color: #4caf50;
    border: none;
    color: white;
    padding: 15px 32px;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 16px;
    margin: 4px 2px;
    cursor: pointer;
}
</style>