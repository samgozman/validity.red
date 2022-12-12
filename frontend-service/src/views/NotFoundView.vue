<script setup lang="ts">
import { RouterLink } from "vue-router";
import StandardContainer from "@/components/layout/StandardContainer.vue";
</script>

<template>
  <StandardContainer>
    <div class="hero min-h-[60vh] bg-base-200">
      <div class="hero-content text-center">
        <div class="w-full">
          <h1 class="text-5xl font-bold">🙄 404 🧐</h1>
          <p class="py-6">The page you are looking for is not found</p>
          <div class="mockup-window border bg-base-300">
            <div class="flex justify-center">
              <img src="/img/travolta404.gif" alt="Travolta not found" />
            </div>
          </div>
          <RouterLink to="/" class="btn btn-primary mt-4 w-full"
            >Home page</RouterLink
          >
        </div>
      </div>
    </div>
  </StandardContainer>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  created() {
    document.title = "Page not found | Validity.Red";
  },
});
</script>
