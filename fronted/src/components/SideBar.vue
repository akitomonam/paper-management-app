<template>
    <div class="sidebar" v-bind:class="{ open: isOpen }">
        <ul>
            <li><router-link to="/">ホーム</router-link></li>
            <li><router-link to="/mypage">マイページ</router-link></li>
            <li><router-link to="/login">ログイン</router-link></li>
            <li><router-link to="/logout">ログアウト</router-link></li>
            <li><router-link to="/SignUplist">ユーザーリスト</router-link></li>
        </ul>
    </div>
</template>

<script>
export default {
    name: "SideBar",
    props: {
        // App.vueからisOpenを受け取るプロパティを定義
        isOpen: {
            type: Boolean,
            required: true,
        },
    },
    methods: {
        changeOpen() {
            // App.vueにchangeOpenイベントを発行する
            this.$emit("changeOpen");
        },
    },
}
</script>

<style>
.sidebar {
    /* サイドバーのスタイルをここに記述 */
    width: 200px;
    /* サイドバーの幅を200pxに指定 */
    position: fixed;
    /* サイドバーを固定表示にする */
    top: 0;
    /* サイドバーをページの上部に配置する */
    bottom: 0;
    /* サイドバーをページの下部まで伸ばす */
    right: 0;
    /* サイドバーをページの右側に配置する */
    background-color: #f5f5f5;
    /* サイドバーの背景色を指定 */
    z-index: 1;
    /* サイドバーを前面に表示する */
    /* サイドバーを非表示にするためのスタイルを追加 */
    transform: translateX(200px);
    transition: transform 0.3s;
}

.sidebar.open {
    /* openクラスが追加された場合、サイドバーを表示する */
    transform: translateX(0);
}

.sidebar ul {
    /* リストのスタイルをここに記述 */
    list-style: none;
    /* リストのマーカーを非表示にする */
    margin: 0;
    /* リストのマージンを0にする */
    padding: 0;
    /* リストのパディングを0にする */
}

.sidebar li {
    /* リストアイテムのスタイルをここに記述 */
    border-bottom: 1px solid #ddd;
    /* リストアイテムの間に区切り線を表示する */
}

.sidebar a {
    /* リンクのスタイルをここに記述 */
    display: block;
    /* リンクをブロック要素として表示する */
    padding: 8px;
    /* リンクのパディングを8pxにする */
    text-decoration: none;
    /* リンクの下線を非表示にする */
    color: #333;
    /* リンクの文字色を#333にする */
    font-size: 14px;
    /* リンクの文字サイズを14pxにする */
}

.sidebar a:hover {
    /* リンクをホバーしたときのスタイルをここに記述 */
    background-color: #eee;
    /* リンクをホバーしたときの背景色を#eeeにする */
    cursor: pointer;
    /* マウスカーソルをポインターにする */
}
</style>
