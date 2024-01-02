import { createRouter, createWebHistory } from "vue-router";
import Home from "../pages/home.vue";
import Test from "../pages/test.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: Home,
    },
    {
      path: "/test",
      name: "test",
      component: Test,
    },
  ],
});

export default router;
