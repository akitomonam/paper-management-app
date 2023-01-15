<template>
    <div>
        <form @submit.prevent="uploadFile">
            <input type="file" ref="fileInput" @change="updateButtonDisabledStatus"/>
            <button type="submit" :disabled="isButtonDisabled">Upload</button>
        </form>
        <vue-element-loading :active="isUploading" is-full-screen text="Now uploading..." size="128"/>
    </div>
</template>

<script>
import axios from 'axios';
import { config } from '../../config';
import VueElementLoading from "vue-element-loading";

export default {
    name: 'UploadForm',
    components: {
        VueElementLoading
    },
    data() {
        return {
            isButtonDisabled: true,
            isUploading: false,
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
            this.isUploading = true; // アップロード開始時に、アップロード中のステータスを表すデータプロパティを true に設定する
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
                    // alert("ファイルアップロード完了\n" + "ファイル名:" + response.data.filename)
                    this.$refs.fileInput.value = ''; // ファイル入力欄をリセット
                    this.isButtonDisabled = true; // アップロード完了後はボタンを disabled 状態にする
                    this.isUploading = false; // アップロード完了後に、アップロード中のステータスを表すデータプロパティを false に設定する
                    // this.$emit('update-upload-status', response.data) // レスポンスボディのfilenameを参照して、responseMessageに代入する
                    this.$emit('update-upload-status') // レスポンスボディのfilenameを参照して、responseMessageに代入する
                } else {
                    this.isUploading = false; // アップロード完了後に、アップロード中のステータスを表すデータプロパティを false に設定する
                    console.error('レスポンスボディが存在しません')
                    alert("アップロードエラーが発生しました")
                }
            } catch (error) {
                this.isUploading = false; // アップロード完了後に、アップロード中のステータスを表すデータプロパティを false に設定する
                alert("アップロードエラーが発生しました")
                console.log('File uploaded Failed');
                console.error(error);
            }
        },
    }
}
</script>

<style scoped>
/* アップロードフォームを装飾する */
form {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
}

/* .uploading-status {
    color: rgb(255, 0, 0);
} */

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

/* .loader,
.loader:before,
.loader:after {
    background: #ffffff;
    -webkit-animation: load1 1s infinite ease-in-out;
    animation: load1 1s infinite ease-in-out;
    width: 1em;
    height: 4em;
}

.loader {
    color: #299ace;
    text-indent: -9999em;
    margin: 88px auto;
    position: relative;
    font-size: 11px;
    -webkit-transform: translateZ(0);
    -ms-transform: translateZ(0);
    transform: translateZ(0);
    -webkit-animation-delay: -0.16s;
    animation-delay: -0.16s;
}

.loader:before,
.loader:after {
    position: absolute;
    top: 0;
    content: '';
}

.loader:before {
    left: -1.5em;
    -webkit-animation-delay: -0.32s;
    animation-delay: -0.32s;
}

.loader:after {
    left: 1.5em;
}

@-webkit-keyframes load1 {

    0%,
    80%,
    100% {
        box-shadow: 0 0;
        height: 4em;
    }

    40% {
        box-shadow: 0 -2em;
        height: 5em;
    }
}

@keyframes load1 {

    0%,
    80%,
    100% {
        box-shadow: 0 0;
        height: 4em;
    }

    40% {
        box-shadow: 0 -2em;
        height: 5em;
    }
} */
</style>
