import { createRouter, createWebHistory } from "vue-router";
import PaperManagement from "../views/HomePage.vue";

const routes = [
    {
        path: "/",
        name: "home",
        component: PaperManagement,
    },
    {
        path: "/login",
        name: "LoginView",
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () =>
            import(/* webpackChunkName: "about" */ "../views/LoginView"),
    },
    {
        path: "/logout",
        name: "LogOut",
        component: () =>
            import(/* webpackChunkName: "about" */ "../views/LogoutView"),
    },
    {
        path: "/mypage",
        name: "MyPage",
        component: () =>
            import(/* webpackChunkName: "about" */ "../views/MyPage"),
    },
    {
        path: "/signup",
        name: "SignUp",
        component: () =>
            import(/* webpackChunkName: "about" */ "../views/SignUp"),
    },
    {
        path: "/SignUplist",
        name: "SignUpList",
        component: () =>
            import(/* webpackChunkName: "about" */ "../views/SignUpList"),
    },
];

const router = createRouter({
    // history: createWebHistory(process.env.BASE_URL),
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

export default router;
