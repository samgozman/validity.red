<script setup lang="ts">
import { RouterView } from "vue-router";
import {
  navbarItems,
  navbarItemsUser,
  navbarItemsUserAuth,
} from "@/components/navbar/items";
import { state } from "@/state";
import NavBar from "@/components/navbar/NavBar.vue";
import NavItem from "@/components/navbar/NavItem.vue";
import Footer from "@/components/FooterComponent.vue";
</script>

<template>
  <div class="drawer h-screen-safe bg-base-200">
    <input id="left-sidebar" type="checkbox" class="drawer-toggle" />
    <div class="drawer-content flex flex-col">
      <header class="bg-base-100">
        <div class="xl:container xl:mx-auto">
          <NavBar />
        </div>
      </header>
      <RouterView />
      <Footer />
    </div>
    <div class="drawer-side max-h-screen-safe">
      <label for="left-sidebar" class="drawer-overlay"></label>
      <div class="menu p-4 overflow-y-auto w-80 bg-base-100">
        <ul>
          <NavItem
            v-for="item in navbarItems.filter((item) => {
              return state.user.isAuthenticated ? item.name !== 'Home' : true;
            })"
            v-bind:key="item.name"
            :item="item"
            :isMobile="true"
          />
        </ul>
        <div class="divider">User section</div>
        <ul>
          <NavItem
            v-for="item in state.user.isAuthenticated
              ? navbarItemsUserAuth
              : navbarItemsUser"
            v-bind:key="item.name"
            :item="item"
            :isMobile="true"
          />
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { RefreshToken } from "@/services/RefreshToken";
import { setUser } from "@/state";

export default defineComponent({
  mounted() {
    // if user has token and user object in local storage then update user state
    const user = localStorage.getItem("user");
    if (RefreshToken.getCookie("token") && user) {
      setUser(JSON.parse(user));
    }
    // Run token refresh task in background
    setInterval(async () => {
      await RefreshToken.call();
    }, 120_000); // every 2 minutes
  },
});
</script>
