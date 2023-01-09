<template>
    <div>
        <!-- 編集用のエリア -->
        <div v-if="editMode">
            <h1>Title: <input v-model="paper.title" type="text"></h1>
            <table>
                <tr>
                    <th>Abstract:</th>
                    <td><input v-model="paper.abstract" type="text"></td>
                </tr>
                <tr>
                    <th>Author:</th>
                    <td><input v-model="paper.author" type="text"></td>
                </tr>
                <tr>
                    <th>Publisher:</th>
                    <td><input v-model="paper.publisher" type="text"></td>
                </tr>
                <tr>
                    <th>Year:</th>
                    <td><input v-model="paper.year" type="number"></td>
                </tr>
            </table>
        </div>
        <!-- 編集前の表示 -->
        <div v-else>
            <div style="display:flex;justify-content:center;">
                <star-rating @update:rating="setRating" v-bind:rating="rating"
                    :max-rating="1" :show-rating="false" :clearable="true"
                    :animate="true" :rounded-corners="true">
                </star-rating>
                <h1 style="margin-left: 10px;">Title:{{ paper.title }}</h1>
            </div>
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
        </div>
        <!-- <button class="edit-button" @click="editPaper">Edit</button> -->
        <button class="edit-button" @click="editMode = !editMode">Edit</button>
        <button v-if="editMode" class="edit-complete-button" @click="editPaper">Complete</button>
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
import StarRating from 'vue-star-rating'
import axios from 'axios'
import { config } from "../../config";
export default {
    components: {
        StarRating
    },
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
            },
            editMode: false,
            rating: 0
        }
    },
    created() {
        axios.get(`${config.URL}:${config.PORT}/api/papers/${this.id}`)
            .then(response => {
                this.paper = response.data
                console.log(this.paper)
                axios.get(`${config.URL}:${config.PORT}/api/checkFavorite`,
                    {
                        params: {
                            sessionToken: localStorage.getItem('sessionToken'),
                            paperId: this.paper.ID,
                        }
                    })
                    .then(response => {
                        console.log(response.data)
                        this.rating = response.data.rating
                    })
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
                            // alert("ファイル削除に成功しました。")
                            this.$emit('update-tables') // 親コンポーネントに発火
                            this.$router.back()
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
        },
        // 編集処理
        async editPaper() {
            // 画面上の文字列を取得する
            let title = this.paper.title
            let abstract = this.paper.abstract
            let author = this.paper.author
            let publisher = this.paper.publisher
            let year = this.paper.year

            // データベースに反映する
            await axios
                .post(
                    `${config.URL}:${config.PORT}/api/editpaperinfo`,
                    {
                        id: this.paper.ID,
                        title: title,
                        abstract: abstract,
                        author: author,
                        publisher: publisher,
                        year: year
                    },
                    {
                        headers: {
                            'Content-Type': 'application/x-www-form-urlencoded'
                        }
                    }
                )
                .then((response) => {
                    console.log("edit response", response);
                    console.log("編集が完了しました");
                    // 画面上の文字列を更新する
                    this.paper.title = response.data.title;
                    this.paper.abstract = response.data.abstract;
                    this.paper.author = response.data.author;
                    this.paper.publisher = response.data.publisher;
                    this.paper.year = response.data.year;
                    this.editMode = false;
                })
                .catch((error) => {
                    console.log("編集APIでエラーが発生しました");
                    console.error(error);
                    alert("編集APIでエラーが発生しました")
                });
        },
        async setRating(rating) {
            this.rating = rating;
            await axios
                .post(
                    `${config.URL}:${config.PORT}/api/favorite`,
                    {
                        sessionToken: localStorage.getItem('sessionToken'),
                        paperId: this.paper.ID,
                        rating: this.rating,
                    },
                    {
                        headers: {
                            'Content-Type': 'application/x-www-form-urlencoded'
                        }
                    }
                )
                .then((response) => {
                    console.log("favorite response", response);
                    console.log("お気に入り編集が完了しました");
                })
                .catch((error) => {
                    console.log("お気に入りAPIでエラーが発生しました");
                    console.error(error);
                    alert("お気に入りAPIでエラーが発生しました")
                });
        },
    },
}
</script>

<style scoped>
table {
    border-collapse: collapse;
    margin: 0 auto;
}
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
