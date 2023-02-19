<template>
  <div>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>Registered Users List</span>
          <!-- <el-button class="button" text>Operation button</el-button> -->
        </div>
      </template>
      <div class="demo-type">
        <div v-for="table in tables" :key="table">
          <el-badge :value="table.PaperCount" class="item" type="primary">
            <el-avatar> {{ table.Username }} </el-avatar>
          </el-badge>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script>
import axios from "axios";
import { config } from "../../config";
export default {
  name: "FileTable",
  data() {
    return {
      tables: [],
    };
  },
  created() {
    this.getUserList();
  },
  computed: {
    // 数値に文字列を追加する
    paperCountStr(cnt) {
      return cnt + " papers";
    },
  },
  methods: {
    getUserList: function () {
      axios.get(`${config.URL}:${config.PORT}/api/userlist`).then((res) => {
        console.log("res.data:", res.data);
        this.tables = res.data;
      });
    },
  },
};
</script>

<style scoped>
.demo-type {
  display: flex;
}

.demo-type > div {
  flex: 1;
  text-align: center;
}

.demo-type > div:not(:last-child) {
  border-right: 1px solid var(--el-border-color);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.box-card {
  width: auto;
  margin: 40px;
}
</style>
