<script setup lang="ts">
import { RouterView } from "vue-router";
import NavBar from "./components/NavBar.vue";
import Footer from "./components/FooterComponent.vue";
</script>

<template>
  <div class="flex flex-col h-screen justify-between">
    <header class="h-auto">
      <div class="container xl:mx-auto">
        <NavBar />
      </div>
    </header>
    <div class="container xl:mx-auto mb-auto">
      <div class="drawer drawer-end">
        <input id="right-sidebar" type="checkbox" class="drawer-toggle" />
        <div class="drawer-content p-6">
          <RouterView />
        </div>
        <div class="drawer-side">
          <label for="right-sidebar" class="drawer-overlay"></label>
          <div class="p-4 overflow-y-auto w-80 bg-base-100 text-base-content">
            Place for login / register form
          </div>
        </div>
      </div>
    </div>
    <Footer />
  </div>
</template>
