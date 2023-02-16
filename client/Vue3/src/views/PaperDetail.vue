<template>
  <div style="text-align: center">
    <!-- 編集用のエリア -->
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>Paper infomation</span>
          <el-button v-if="editMode" class="button" text @click="editAutoPaper"
            ><el-icon> <MagicStick /> </el-icon>Auto</el-button
          >
          <el-button v-if="editMode" class="button" text @click="editPaper"
            ><el-icon el-icon--left> <Finished /> </el-icon>Comlete</el-button
          >
          <el-button
            v-if="!editMode"
            class="button"
            text
            @click="showFile(paper.ID)"
            ><el-icon el-icon--left> <View /> </el-icon>Preview</el-button
          >
          <el-button
            v-if="!editMode"
            class="button"
            text
            @click="editMode = !editMode"
            ><el-icon el-icon--left> <Edit /> </el-icon>Edit</el-button
          >
        </div>
      </template>
      <div v-if="editMode">
        <el-form
          :label-position="labelPosition"
          label-width="100px"
          :model="formLabelAlign"
          style="max-width: 460px  text-align: center;"
        >
          <el-form-item label="Title">
            <el-input v-model="paper.title" />
          </el-form-item>
          <el-form-item label="Abstract">
            <el-input v-model="paper.abstract" />
          </el-form-item>
          <el-form-item label="Author">
            <el-input v-model="paper.author" />
          </el-form-item>
          <el-form-item label="Publisher">
            <el-input v-model="paper.publisher" />
          </el-form-item>
          <el-form-item label="Year">
            <el-input-number
              v-model="paper.year"
              class="mx-4"
              :min="1"
              controls-position="right"
            />
          </el-form-item>
          <el-form-item label="Keywords">
            <template v-for="(item, index) in keywords">
              <template v-if="item.Flag">
                <el-tag
                  :key="index"
                  class="mx-1"
                  closable
                  :disable-transitions="false"
                  @close="handleClose(index)"
                  style="margin-right: 5px"
                >
                  {{ item.Keyword }}
                </el-tag>
              </template>
            </template>
            <el-input
              v-if="inputVisible"
              ref="InputRef"
              v-model="inputValue"
              class="ml-1 w-20"
              size="small"
              @keyup.enter="handleInputConfirm"
              @blur="handleInputConfirm"
            />
            <el-button
              v-else
              class="button-new-tag ml-1"
              size="small"
              @click="showInput"
            >
              + New Keyword
            </el-button>
          </el-form-item>
        </el-form>
      </div>
      <!-- 編集前の表示 -->
      <div v-else>
        <div class="rate-block">
          <el-rate
            @change="setRating"
            v-model="rating"
            :colors="['#99A9BF', '#F7BA2A', '#FF9900']"
          />
        </div>
        <el-descriptions
          :title="paper.title"
          column="1"
          size="default"
          direction="horizontal"
          border
        >
          <el-descriptions-item v-if="paper.title == ``" label="Title">
            {{ paper.title }}
          </el-descriptions-item>
          <el-descriptions-item label="Abstract">{{
            paper.abstract
          }}</el-descriptions-item>
          <el-descriptions-item label="Author">{{
            paper.author
          }}</el-descriptions-item>
          <el-descriptions-item label="Publisher">{{
            paper.publisher
          }}</el-descriptions-item>
          <el-descriptions-item label="Year">{{
            paper.year
          }}</el-descriptions-item>
          <el-descriptions-item label="Keywords">
            <el-tag
              v-for="item in filteredKeywords"
              :key="item"
              class="mx-1"
              :disable-transitions="false"
              style="margin-right: 5px"
            >
              {{ item.Keyword }}
            </el-tag>
          </el-descriptions-item>
        </el-descriptions>
      </div>
      <el-button
        v-if="editMode"
        size="small"
        type="danger"
        bg
        text
        @click="deleteFile(paper.ID)"
        ><el-icon el-icon--left> <Delete /> </el-icon>Delete</el-button
      >
    </el-card>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>File infomation</span>
        </div>
      </template>
      <el-descriptions column="1" size="default" direction="horizontal" border>
        <el-descriptions-item label="ID">{{ paper.ID }}</el-descriptions-item>
        <el-descriptions-item label="File name">{{
          paper.file_name
        }}</el-descriptions-item>
        <el-descriptions-item label="Uploader's ID">{{
          paper.user_id
        }}</el-descriptions-item>
        <el-descriptions-item label="Created at">{{
          paper.created_at
        }}</el-descriptions-item>
        <el-descriptions-item label="Support files">
          <el-button-group
            v-for="item in supportFiles"
            :key="item"
            style="margin-right: 5px"
          >
            <el-button
              class="button"
              type="info"
              plain
              @click="showSupportFile(item.ID)"
              ><el-icon el-icon--left> <View /> </el-icon>{{ item.file_name }}
            </el-button>
            <el-button plain type="info" @click="deleteSupportFile(item.ID)"
              ><el-icon><Close /></el-icon
            ></el-button>
          </el-button-group>
        </el-descriptions-item>
      </el-descriptions>
      <UploadForm :paperId="paper.ID" @update-upload-status="getDB" />
    </el-card>
    <br />
    <!-- <button class="get-bibtex-button" @click="getBibTeX(paper.ID)">GetBibTeX</button> -->
    <CommentComponent :paper_id="paper.ID" />
    <vue-element-loading
      :active="isLoading"
      is-full-screen
      text="Now loading..."
      size="128"
    />
  </div>
</template>

<script>
import VueElementLoading from "vue-element-loading";
import axios from "axios";
import CommentComponent from "../components/CommentComponent.vue";
import UploadForm from "../components/UploadForm.vue";
import { config } from "../../config";
export default {
  components: {
    CommentComponent,
    VueElementLoading,
    UploadForm,
  },
  props: ["id"],
  data() {
    return {
      paper: {
        ID: "",
        abstract: "",
        author: "",
        created_at: "",
        file_name: "",
        file_path: "",
        publisher: "",
        title: "",
        user_id: "",
        year: "",
        bibtex: "",
      },
      keywords: [],
      supportFiles: ["hoge", "fuga", "piyo"],
      inputValue: "",
      inputVisible: false,
      editMode: false,
      rating: 0,
      isLoading: false,
    };
  },
  computed: {
    filteredKeywords() {
      return this.keywords.filter((item) => item.Flag);
    },
  },
  created() {
    axios
      .get(`${config.URL}:${config.PORT}/api/papers/${this.id}`)
      .then((response) => {
        this.paper = response.data.Paper;
        this.keywords = response.data.Keywords;
        axios
          .get(`${config.URL}:${config.PORT}/api/checkFavorite`, {
            params: {
              sessionToken: localStorage.getItem("sessionToken"),
              paperId: this.paper.ID,
            },
          })
          .then((response) => {
            console.log(response.data);
            this.rating = response.data.rating;
          });
      });
    this.getDB();
  },
  methods: {
    // tableをクリックした際に実行される処理
    async showFile(targetFileId) {
      console.log("click table!!!");
      console.log("targetFileId:", targetFileId); // 目的の文字列を出力する
      // サーバーに保管されているファイルをプレビューする
      axios
        .get(`${config.URL}:${config.PORT}/api/preview?fileId=${targetFileId}`)
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
    async showSupportFile(targetFileId) {
      console.log("click table!!!");
      console.log("targetFileId:", targetFileId); // 目的の文字列を出力する
      // サーバーに保管されているファイルをプレビューする
      axios
        .get(
          `${config.URL}:${config.PORT}/api/previewSupportFile?fileId=${targetFileId}`
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
      console.log("targetFileId:", targetFileId);
      if (confirm("本当に削除しますか？")) {
        try {
          const response = await axios.get(
            `${config.URL}:${config.PORT}/api/delete?fileId=${targetFileId}`
          );
          if (response.data) {
            // レスポンスボディが存在する場合
            console.log("response.data", response.data);
            if (response.data.result == "true") {
              // alert("ファイル削除に成功しました。")
              this.$emit("update-tables"); // 親コンポーネントに発火
              this.$router.back();
            } else {
              alert("ファイル削除に失敗しました。");
            }
          } else {
            console.error("レスポンスボディが存在しません");
          }
        } catch {
          console.log("File delete Failed");
        }
      }
    },
    async deleteSupportFile(targetFileId) {
      console.log("targetFileId:", targetFileId);
      if (confirm("本当に削除しますか？")) {
        try {
          const response = await axios.get(
            `${config.URL}:${config.PORT}/api/deleteSupportFile?fileId=${targetFileId}`
          );
          if (response.data) {
            // レスポンスボディが存在する場合
            console.log("response.data", response.data);
            if (response.data.result == "true") {
              // alert("ファイル削除に成功しました。")
              this.$emit("update-tables"); // 親コンポーネントに発火
              this.getDB();
            } else {
              alert("ファイル削除に失敗しました。");
            }
          } else {
            console.error("レスポンスボディが存在しません");
          }
        } catch {
          console.log("File delete Failed");
        }
      }
    },
    // 編集処理
    async editPaper() {
      // 画面上の文字列を取得する
      let title = this.paper.title;
      let abstract = this.paper.abstract;
      let author = this.paper.author;
      let publisher = this.paper.publisher;
      let year = this.paper.year;
      let keywords = this.keywords;
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
            year: year,
            keywords: keywords,
          },
          {
            headers: {
              "Content-Type": "application/json",
            },
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
          alert("編集APIでエラーが発生しました");
        });
    },
    // 自動メタ情報付与
    async editAutoPaper() {
      this.isLoading = true;
      await axios
        .post(
          `${config.URL}:${config.PORTPYTHON}/api/python/get_paper_meta_data`,
          {
            file_path: this.paper.file_path,
          },
          {
            headers: {
              "Content-Type": "application/json",
            },
          }
        )
        .then((response) => {
          console.log("edit auto response", response);
          console.log("自動補完が完了しました");
          // 画面上の文字列を更新する
          this.paper.title = response.data.title;
          this.paper.abstract = response.data.abstract;
          this.paper.author = response.data.author;
          this.paper.year = response.data.year;
          // this.paper.publisher = response.data.publisher;
          // this.editMode = false;
        })
        .catch((error) => {
          console.log("自動編集APIでエラーが発生しました");
          console.error(error);
          alert("自動編集APIでエラーが発生しました");
        });
      this.isLoading = false;
    },
    // bibTeX情報取得
    // async getBibTeX(targetFileId) {
    //     this.isLoading = true;
    //     await axios.get(`${config.URL}:${config.PORT}/api/getBibTeX?fileId=${targetFileId}`)
    //         .then((response) => {
    //             console.log("edit auto response", response);
    //             console.log("bibTeX情報取得が完了しました");
    //             this.paper.bibtex = response.data.bibtex
    //         })
    //         .catch((error) => {
    //             console.log("BibTex取得APIでエラーが発生しました");
    //             console.error(error);
    //             alert("BibTex取得APIでエラーが発生しました")
    //         });
    //     this.isLoading = false;
    // },
    async setRating() {
      await axios
        .post(
          `${config.URL}:${config.PORT}/api/favorite`,
          {
            sessionToken: localStorage.getItem("sessionToken"),
            paperId: this.paper.ID,
            rating: this.rating,
          },
          {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded",
            },
          }
        )
        .then((response) => {
          console.log("favorite response", response);
          console.log("お気に入り編集が完了しました");
        })
        .catch((error) => {
          console.log("お気に入りAPIでエラーが発生しました");
          console.error(error);
          alert("お気に入りAPIでエラーが発生しました");
        });
    },
    showInput() {
      this.inputVisible = true;
    },
    handleClose(index) {
      console.log("index:", index);
      this.keywords[index]["Flag"] = false;
      console.log("this.keywords:", this.keywords);
    },
    handleInputConfirm() {
      if (this.inputValue) {
        // ここでkeywordsのキーでinputValueが存在していた場合でflagがfalseの場合はflagをtrueにする
        // inputValueが存在していてflagがtrueの場合はalertを出す
        if (
          this.keywords.some((keyword) => keyword.keyword === this.inputValue)
        ) {
          if (
            this.keywords.some(
              (keyword) =>
                keyword.keyword === this.inputValue && keyword.Flag === false
            )
          ) {
            this.keywords.find(
              (keyword) => keyword.keyword === this.inputValue
            )["Flag"] = true;
          } else {
            alert("すでに存在するキーワードです");
          }
        } else {
          this.keywords.push({ Keyword: this.inputValue, Flag: true });
        }
      }
      console.log("keywords:", this.keywords);
      this.inputVisible = false;
      this.inputValue = "";
    },
    getDB() {
      console.log("paperId:", this.id);
      axios
        .get(`${config.URL}:${config.PORT}/api/supportFiles?paperId=${this.id}`)
        .then((res) => {
          this.supportFiles = res.data;
          console.log("supportFiles:", this.supportFiles);
        });
    },
  },
};
</script>

<style scoped>
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
