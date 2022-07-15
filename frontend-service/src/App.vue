<script setup lang="ts">
import { RouterView } from "vue-router";
import NavBar from "./components/NavBar.vue";
import NavBarItems from "./components/NavBarItems.vue";
import Footer from "./components/FooterComponent.vue";
</script>

<template>
  <div class="drawer">
    <input id="left-sidebar" type="checkbox" class="drawer-toggle" />
    <div class="drawer-content flex flex-col">
      <header>
        <div class="xl:container xl:mx-auto">
          <NavBar />
        </div>
      </header>
      <div class="container sm:mx-auto mb-auto px-6">
        <RouterView />
      </div>
      <Footer />
    </div>
    <div class="drawer-side">
      <label for="left-sidebar" class="drawer-overlay"></label>
      <ul class="menu p-4 overflow-y-auto w-80 bg-base-100">
        <NavBarItems />
      </ul>
    </div>
  </div>
</template>
