<script setup lang="ts">
import { RouterView } from "vue-router";
import NavBar from "./components/NavBar.vue";
import Footer from "./components/FooterComponent.vue";
</script>

<template>
  <div class="flex flex-col h-screen justify-between">
    <header class="h-20">
      <div class="xl:container xl:mx-auto">
        <NavBar />
      </div>
    </header>
    <div class="container xl:mx-auto mb-auto">
      <RouterView />
    </div>
    <Footer />
  </div>
</template>
