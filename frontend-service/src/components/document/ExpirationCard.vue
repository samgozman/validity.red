<script setup lang="ts">
import { defineProps } from "vue";
defineProps<{
  expiresAt: string;
}>();
</script>

<template>
  <div class="card shadow-lg bg-base-100">
    <div class="card-body">
      <div class="grid grid-cols-2 gap-0">
        <h3 class="text-sm font-bold">Valid until:</h3>
        <p class="text-sm text-right">{{ expiresAt }}</p>
      </div>
    </div>
  </div>
</template>
