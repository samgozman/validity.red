<script setup lang="ts">
import { defineProps } from "vue";
import { RouterLink } from "vue-router";
defineProps<{
  title: string;
}>();
</script>

<!-- TODO: Create dynamicly from browser history -->
<template>
  <div class="text-sm breadcrumbs">
    <ul>
      <li>
        <RouterLink to="/documents">Documents</RouterLink>
      </li>
      <li>{{ title }}</li>
    </ul>
  </div>
</template>
