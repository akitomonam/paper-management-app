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
        path: "/mypage",
        name: "MyPage",
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () =>
            import(/* webpackChunkName: "about" */ "../views/MyPage"),
    },
];

const router = createRouter({
    // history: createWebHistory(process.env.BASE_URL),
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

export default router;
