<template>
    <form @submit.prevent="uploadFile">
        <input type="file" ref="fileInput" @change="updateButtonDisabledStatus"/>
        <button type="submit" :disabled="isButtonDisabled">Upload</button>
        <!-- アップロード中のステータスを表す要素を表示する -->
        <template v-if="isUploading">
            <div class="uploading-status">
                Uploading{{ '.'.repeat(uploadingStatusDots) }}
            </div>
        </template>
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
            isButtonDisabled: true,
            isUploading: false,
            uploadingStatusDots: 0,
        };
    },
    created() {
    },
    methods: {
        updateButtonDisabledStatus() {
            this.isButtonDisabled = !this.$refs.fileInput.files[0];
        },
        async uploadFile() {
            const sessionToken = localStorage.getItem('sessionToken');
            console.log("sessionToken:", sessionToken)
            if (!sessionToken) {
                alert("ログイン後にアップロードしてください")
                return
            }

            this.isUploading = true; // アップロード開始時に、アップロード中のステータスを表すデータプロパティを true に設定する
            // アップロード中に "." の数を増やす処理を開始
            this.uploadingStatusIntervalId = setInterval(() => {
                // アップロード中に表示される "." の数を増やす
                this.uploadingStatusDots = (this.uploadingStatusDots + 1) % 4;
            }, 500);
            const file = this.$refs.fileInput.files[0];
            const formData = new FormData();
            formData.append('file', file);
            try {
                const response = await axios.post(`${config.URL}:${config.PORT}/upload/file`, formData, {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                        'Authorization': `Bearer ${sessionToken}`
                    },
                })
                if (response.data) { // レスポンスボディが存在する場合
                    alert("ファイルアップロード完了\n" + "ファイル名:" + response.data.filename)
                    this.$refs.fileInput.value = ''; // ファイル入力欄をリセット
                    this.isButtonDisabled = true; // アップロード完了後はボタンを disabled 状態にする
                    this.isUploading = false; // アップロード完了後に、アップロード中のステータスを表すデータプロパティを false に設定する
                    // this.$emit('update-upload-status', response.data) // レスポンスボディのfilenameを参照して、responseMessageに代入する
                    this.$emit('update-upload-status') // レスポンスボディのfilenameを参照して、responseMessageに代入する
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

.uploading-status {
    color: rgb(255, 0, 0);
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

/* ボタンが disabled 状態の場合のスタイル */
button[type="submit"][disabled] {
    background-color: #9e9e9e;
    cursor: not-allowed;
}
</style>
