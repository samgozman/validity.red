<script setup lang="ts">
import { defineProps } from "vue";
import { RouterLink } from "vue-router";
defineProps<{
  title: string;
}>();
</script>

<!-- TODO: Create dynamically from browser history -->
<template>
  <div class="text-sm breadcrumbs">
    <ul>
      <li>
        <RouterLink to="/documents">
          <ion-icon name="documents-outline" class="mr-1"></ion-icon>
          Documents
        </RouterLink>
      </li>
      <li>
        <ion-icon name="document-text-outline" class="mr-1"></ion-icon>
        {{ title }}
      </li>
    </ul>
  </div>
</template>
