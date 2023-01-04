<template>
    <div>
        <h1>{{ paper.title }}</h1>
        <table>
            <tr>
                <th>Abstract:</th>
                <td>{{ paper.abstract }}</td>
            </tr>
            <tr>
                <th>Author:</th>
                <td>{{ paper.author }}</td>
            </tr>
            <tr>
                <th>Publisher:</th>
                <td>{{ paper.publisher }}</td>
            </tr>
            <tr>
                <th>Year:</th>
                <td>{{ paper.year }}</td>
            </tr>
        </table>
        <h2>File information</h2>
        <table>
            <tr>
                <th>ID:</th>
                <td>{{ paper.ID }}</td>
            </tr>
            <tr>
                <th>File name:</th>
                <td>{{ paper.file_name }}</td>
            </tr>
            <tr>
                <th>Uploader's ID:</th>
                <td>{{ paper.user_id }}</td>
            </tr>
            <tr>
                <th>Created at:</th>
                <td>{{ paper.created_at }}</td>
            </tr>
        </table>
        <br>
        <button class="preview-button" @click="showFile(paper.ID)">Preview</button>
        <button class="delete-button" @click="deleteFile(paper.ID)">Delete</button>
    </div>
</template>

<script>
import axios from 'axios'
import { config } from "../../config";
export default {
    props: ['id'],
    data() {
        return {
            paper: {
                ID: '',
                abstract: '',
                author: '',
                created_at: '',
                file_name: '',
                file_path: '',
                publisher: '',
                title: '',
                user_id: '',
                year: ''
            }
        }
    },
    created() {
        axios.get(`${config.URL}:${config.PORT}/api/papers/${this.id}`)
            .then(response => {
                this.paper = response.data
                console.log(this.paper)
            })
    },
    methods: {
        // tableをクリックした際に実行される処理
        async showFile(targetFileId) {
            console.log("click table!!!");
            console.log("targetFileId", targetFileId); // 目的の文字列を出力する
            // サーバーに保管されているファイルをプレビューする
            axios
                .get(
                    `${config.URL}:${config.PORT}/api/preview?fileId=${targetFileId}`
                )
                .then((response) => {
                    // プレビューするファイルのURLを取得する
                    let filepath = response.data.fileUrl;
                    // プレビューするファイルのURLをもとに、新しいタブを開く
                    filepath = filepath.replace(/^\./, "");
                    window.open(`${config.URL}:${config.PORT}${filepath}`, "_blank");
                })
                .catch((error) => {
                    console.log("ファイルプレビューAPIでエラーが発生しました");
                    console.error(error);
                });
        },
        async deleteFile(targetFileId) {
            console.log("targetFileId:", targetFileId)
            if (confirm('本当に削除しますか？')) {
                try {
                    const response = await axios.get(`${config.URL}:${config.PORT}/api/delete?fileId=${targetFileId}`)
                    if (response.data) { // レスポンスボディが存在する場合
                        console.log("response.data", response.data);
                        if (response.data.result == "true") {
                            alert("ファイル削除に成功しました。")
                            this.$emit('update-tables') // 親コンポーネントに発火
                        } else {
                            alert("ファイル削除に失敗しました。")
                        }
                    } else {
                        console.error('レスポンスボディが存在しません')
                    }
                } catch {
                    console.log('File delete Failed');
                }
            }
        }
    },
}
</script>

<style>
.preview-button {
    color: white;
    background-color: rgb(88, 88, 226);
    height: 50px;
    width: 100px;
    border: none;
    margin-right: 10px;
    cursor: pointer;
}

.delete-button {
    color: white;
    background-color: rgb(239, 58, 58);
    height: 50px;
    width: 100px;
    border: none;
    cursor: pointer;
}
</style>
