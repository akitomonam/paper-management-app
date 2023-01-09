<template>
    <div class="logout">
    <div class="logout-triangle"></div>

    <h2 class="logout-header">Log out</h2>

    <form class="logout-container">
        <p><input type="submit" value="Log out" @click.prevent="logout"></p>
    </form>
</div>
</template>
<script>
export default {
    name: "LogOut",
    methods: {
        logout() {
            // ローカルストレージからセッショントークンを削除する
            localStorage.removeItem('sessionToken');
            // ログイン画面にリダイレクトする
            this.$router.push({ path: '/login' });
        },
    },
}
</script>

<style>
@import url(https://fonts.googleapis.com/css?family=Open+Sans:400,700);

body {
    background: rgba(255, 255, 255, 0);
    font-family: 'Open Sans', sans-serif;
}

.logout {
    width: 400px;
    margin: 16px auto;
    font-size: 16px;
}

/* Reset top and bottom margins from certain elements */
.logout-header,
.logout p {
    margin-top: 0;
    margin-bottom: 0;
}

/* The triangle form is achieved by a CSS hack */
.logout-triangle {
    width: 0;
    margin-right: auto;
    margin-left: auto;
    border: 12px solid transparent;
    border-bottom-color: #28d;
}

.logout-header {
    background: #28d;
    padding: 20px;
    font-size: 1.4em;
    font-weight: normal;
    text-align: center;
    text-transform: uppercase;
    color: #fff;
}

.logout-container {
    background: #ebebeb;
    padding: 12px;
}

/* Every row inside .logout-container is defined with p tags */
.logout p {
    padding: 12px;
}

.logout input {
    box-sizing: border-box;
    display: block;
    width: 100%;
    border-width: 1px;
    border-style: solid;
    padding: 16px;
    outline: 0;
    font-family: inherit;
    font-size: 0.95em;
}

.logout input[type="submit"] {
    background: #28d;
    border-color: transparent;
    color: #fff;
    cursor: pointer;
}

.logout input[type="submit"]:hover {
    background: #17c;
}

/* Buttons' focus effect */
.logout input[type="submit"]:focus {
    border-color: #05a;
}
</style>
