<template>
  <div>
  <img alt="Vue logo" src="./assets/logo.png">
    <form @submit.prevent="uploadFile">
      <input type="file" ref="fileInput" />
      <button type="submit">アップロード</button>
    </form>
    <p>filename:{{ file_name }}</p>
    <p>status:{{ upload_status }}</p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      file_name: '', //Goから受け取るメッセージ
      upload_status: '', //Goから受け取るメッセージ
    };
  },
  methods: {
    async uploadFile() {
      const file = this.$refs.fileInput.files[0];
      const formData = new FormData();
      formData.append('file', file);

      try {
        const response = await axios.post('http://localhost:12345/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        },
          proxy: false //ローカルホストなのでプロキシを経由しない
        })
        if (response.data) { // レスポンスボディが存在する場合
          this.file_name = response.data.filename // レスポンスボディのfilenameを参照して、responseMessageに代入する
          this.upload_status = response.data.status
        } else {
          console.error('レスポンスボディが存在しません')
        }
      } catch (error) {
        console.log('File uploaded Failed');
        console.error(error);
      }
    },
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
