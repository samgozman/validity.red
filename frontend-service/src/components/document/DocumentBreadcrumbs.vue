<script setup lang="ts">
import { defineProps } from "vue";
import { RouterLink } from "vue-router";
import { DocumentOutline, DocumentTextOutline } from "@vicons/ionicons5";
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
          <DocumentOutline class="w-4 mr-1" />
          Documents
        </RouterLink>
      </li>
      <li>
        <DocumentTextOutline class="w-4 mr-1" />
        {{ title }}
      </li>
    </ul>
  </div>
</template>
