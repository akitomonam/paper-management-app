import { createRouter, createWebHistory } from "vue-router";
import PaperManagement from "../views/HomePage.vue";

const routes = [
    {
        path: "/",
        name: "home",
        component: PaperManagement,
        meta: { requiresAuth: true },
    },
    {
        path: "/login",
        name: "LoginView",
        component: () =>
            import(/* webpackChunkName: "about" */ "../views/LoginView"),
        beforeEnter(to, from, next) {
            // ここでセッショントークンを確認する
            const sessionToken = localStorage.getItem('sessionToken')
            console.log("sessionToken:", !!sessionToken)
            if (sessionToken) {
                alert("既にログインしています")
                from()
            } else {
                next()
            }
        }
    },
    {
        path: "/logout",
        name: "LogOut",
        component: () =>
            import(/* webpackChunkName: "about" */ "../views/LogoutView"),
        meta: { requiresAuth: true },
    },
    {
        path: "/mypage",
        name: "MyPage",
        component: () => import(/* webpackChunkName: "about" */ "../views/MyPage"),
        meta: { requiresAuth: true },
    },
    {
        path: "/signup",
        name: "SignUp",
        component: () => import(/* webpackChunkName: "about" */ "../views/SignUp"),
    },
    {
        path: "/SignUplist",
        name: "SignUpList",
        component: () =>
            import(/* webpackChunkName: "about" */ "../views/SignUpList"),
    },
    {
        path: "/papers/:id",
        name: "PaperDetail",
        props: true,
        component: () =>
            import(/* webpackChunkName: "about" */ "../views/PaperDetail"),
        meta: { requiresAuth: true },
    },
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

router.beforeEach((to) => {
    const sessionToken = localStorage.getItem('sessionToken')
    if (to.meta.requiresAuth && sessionToken == null) {
        alert("ログインしてください")
        return { name: "LoginView" };
    }
});
export default router;
