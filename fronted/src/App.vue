<template>
  <div>
    <button class="settings-button" @click="changeOpen">
      &#9776;
    </button>
    <HeaderComponent />
    <SideBar v-bind:isOpen="isOpen" @changeOpen="changeOpen" />
    <router-view />
  </div>
</template>

<script>
import HeaderComponent from "./components/HeaderComponent.vue";
import SideBar from "./components/SideBar.vue";

export default {
  name: "PaperManagement",
  components: {
    HeaderComponent,
    SideBar,
  },
  data() {
    return {
      isOpen: false, // サイドバーの表示状態を管理するデータプロパティ
    }
  },
  methods: {
    changeOpen() {
      // サイドバーの表示状態を反転する
      this.isOpen = !this.isOpen;
    },
  },
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

.settings-button {
  position: absolute;
  /* 設定ボタンを絶対配置にする */
  top: 0px;
  /* 設定ボタンをページの上部に配置する */
  right: 10px;
  /* 設定ボタンをページの右上に配置する */
  background-color: #fff;
  /* 設定ボタンの背景色を指定 */
  border: none;
  /* 設定ボタンの枠線を非表示にする */
  cursor: pointer;
  /* マウスカーソルをポインターにする */
  font-size: 20px;
  /* 設定ボタンの文字サイズを指定 */
  z-index: 2;
  /* 設定ボタンを前面に表示する */
}

.settings-button:hover {
  /* 設定ボタンにマウスが乗った時の装飾を追加 */
  background-color: #eee;
  /* 設定ボタンの背景色を指定 */
}
</style>
