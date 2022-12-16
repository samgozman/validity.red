<script setup lang="ts">
import type { RouteLocationRaw } from "vue-router";
import { RouterLink } from "vue-router";
import { defineProps } from "vue";

defineProps<{
  location: RouteLocationRaw;
  name: string;
}>();
</script>

<template>
  <RouterLink
    :to="location"
    class="btn border-none md:px-10 bg-base-100 text-primary hover:shadow-lg hover:bg-primary hover:text-base-100 font-medium transition duration-300 ease-in-out"
  >
    {{ name }}
  </RouterLink>
</template>
