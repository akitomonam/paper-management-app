<template>
  <div>
    <el-upload
      class="upload-demo"
      drag
      :action="uploadUrl"
      :on-success="handleSuccess"
      :on-error="handleError"
      multiple
      :headers="authHeaders"
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">
        Drop file here or <em>click to upload</em>
      </div>
      <template #tip>
        <div class="el-upload__tip">
          pdf/pptx files with a size less than 100MB
        </div>
      </template>
    </el-upload>
  </div>
</template>

<script>
import { config } from "../../config";
import { ElNotification } from "element-plus";

export default {
  name: "UploadForm",
  components: {},
  data() {
    return {
      uploadUrl: `${config.URL}:${config.PORT}/upload/file`,
    };
  },
  created() {},
  computed: {
    authHeaders() {
      const token = localStorage.getItem("sessionToken");
      return { Authorization: "Bearer " + token };
    },
  },
  methods: {
    handleSuccess() {
      // アップロード成功時の処理を記述する
      ElNotification({
        title: "Success",
        message: "Upload succeeded",
        type: "success",
      });
      this.$emit("update-upload-status"); // レスポンスボディのfilenameを参照して、responseMessageに代入する
    },
    handleError(error) {
      // エラー時の処理を記述する
      ElNotification({
        title: "Error",
        message: `Error: ${error.message}`,
        type: "error",
      });
    },
  },
};
</script>

<style scoped>
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

/* ボタンが disabled 状態の場合のスタイル */
button[type="submit"][disabled] {
  background-color: #9e9e9e;
  cursor: not-allowed;
}
</style>
