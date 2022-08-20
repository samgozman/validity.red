<script setup lang="ts">
import { RouterView } from "vue-router";
import { navbarItems } from "@/components/navbar/items";
import NavBar from "@/components/navbar/NavBar.vue";
import NavBarItems from "@/components/navbar/NavBarItems.vue";
import Footer from "@/components/FooterComponent.vue";
</script>

<template>
  <div class="drawer h-screen bg-base-200">
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
    <div class="drawer-side">
      <label for="left-sidebar" class="drawer-overlay"></label>
      <ul class="menu p-4 overflow-y-auto w-80 bg-base-100">
        <NavBarItems :items="navbarItems" />
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { RefreshToken } from "./services/RefreshToken";

export default defineComponent({
  mounted() {
    // Run token refresh task in background
    setInterval(async () => {
      await RefreshToken.call();
    }, 30000);
  },
});
</script>
