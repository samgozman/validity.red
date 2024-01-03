import { createRouter, createWebHistory } from "vue-router";
import MainView from "@/views/MainView.vue";
import AboutView from "@/views/AboutView.vue";
import LoginView from "@/views/LoginView.vue";
import EmailConfirmViewVue from "@/views/EmailConfirmView.vue";
import { logout } from "@/services/Logout";

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
      component: () => import("@/views/RegistrationView.vue"),
    },
    {
      path: "/verify",
      name: "verify",
      component: EmailConfirmViewVue,
    },
    {
      path: "/logout",
      name: "logout",
      redirect: () => {
        logout();
        return { name: "home" };
      },
    },
    {
      path: "/dashboard",
      name: "dashboard",
      // route level code-splitting
      // this generates a separate chunk (Dashboard.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("@/views/DashboardView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/documents",
      name: "documents",
      component: () => import("@/views/DocumentsListView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/documents/create",
      name: "document-create",
      component: () => import("@/views/DocumentCreateView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/documents/:id",
      name: "document",
      component: () => import("@/views/DocumentView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/documents/:id/edit",
      name: "document-edit",
      component: () => import("@/views/DocumentCreateView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/about",
      name: "about",
      component: AboutView,
    },
    // Should be the last route
    {
      path: "/:pathMatch(.*)",
      name: "not-found",
      component: () => import("@/views/NotFoundView.vue"),
    },
  ],
});

router.beforeEach(async (to) => {
  console.log("auth guard");
  // Check for requiresAuth guard
  if (!to.meta.requiresAuth) return;

  console.log("cookies:", document.cookie);

  // check if cookie is exists
  const cookie = document.cookie.match(
    new RegExp("(^| )" + "token" + "=([^;]+)")
  );
  console.log("cookie matched:", cookie);

  // TODO: check that user object is exists as well

  if (
    !cookie &&
    // Avoid an infinite redirect
    to.name !== "login"
  ) {
    console.log("redirect to login");
    // redirect the user to the login page
    return { name: "login", query: { redirect: to.fullPath } };
  }
});

export default router;
