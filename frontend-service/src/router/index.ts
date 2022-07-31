import { createRouter, createWebHistory } from "vue-router";
import MainView from "../views/MainView.vue";
import AboutView from "../views/AboutView.vue";
import LoginView from "../views/LoginView.vue";
import RegistrationView from "../views/RegistrationView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: MainView,
    },
    {
      path: "/login",
      name: "login",
      component: LoginView,
    },
    {
      path: "/register",
      name: "register",
      component: RegistrationView,
    },
    {
      path: "/dashboard",
      name: "dashboard",
      // route level code-splitting
      // this generates a separate chunk (Dashboard.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/DashboardView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/documents",
      name: "documents",
      // route level code-splitting
      // this generates a separate chunk (Documents.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/DocumentsListView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/documents/:id",
      name: "document",
      // route level code-splitting
      // this generates a separate chunk (Documents.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/DocumentView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/about",
      name: "about",
      component: AboutView,
    },
  ],
});

router.beforeEach(async (to, from) => {
  // Check for requiresAuth guard
  if (!to.meta.requiresAuth) return;

  // check if cookie is exists
  const cookie = document.cookie.match(
    new RegExp("(^| )" + "token" + "=([^;]+)")
  );

  // TODO: check that user object is exists as well

  if (
    !cookie &&
    // Avoid an infinite redirect
    to.name !== "login"
  ) {
    // redirect the user to the login page
    return { name: "login" };
  }
});

export default router;
