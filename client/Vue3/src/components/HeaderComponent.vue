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
                            <el-dropdown-item><el-link @click="$router.push({ path: '/mypage' })">My Page</el-link></el-dropdown-item>
                            <el-dropdown-item><el-link @click="$router.push({ path: '/login' })">Login</el-link></el-dropdown-item>
                            <el-dropdown-item><el-link @click="$router.push({ path: '/logout' })">Logout</el-link></el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </div>
        </el-header>
        <hr>
        <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
            <el-menu-item index="1"><el-link @click="$router.push({ path: '/' })">Home</el-link></el-menu-item>
            <el-menu-item index="2"><el-link @click="$router.push({ path: '/SignUplist' })">User List</el-link></el-menu-item>
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
