<template>
    <div class="common-layout">
        <el-header class="custom-header" direction="horizontal">
            <el-avatar shape="square" src="https://avatars.githubusercontent.com/u/108845832?s=200&v=4" />
            <span class="title">Paper Management App</span>
            <div class="toolbar">
                <el-dropdown>
                    <el-avatar> user </el-avatar>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <el-dropdown-item><router-link to="/mypage">My Page</router-link></el-dropdown-item>
                            <el-dropdown-item><router-link to="/login">Login</router-link></el-dropdown-item>
                            <el-dropdown-item><router-link to="/logout">Logout</router-link></el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </div>
        </el-header>
        <hr>
        <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
            <el-menu-item index="1"><router-link to="/">Home</router-link></el-menu-item>
            <el-menu-item index="2"><router-link to="/SignUplist">User List</router-link></el-menu-item>
        </el-menu>
    </div>
</template>

<style scoped>
.custom-header {
    height: auto;
    padding: 0;
    display: flex;
    align-items: center;
    justify-content: space-between;
}
</style>
