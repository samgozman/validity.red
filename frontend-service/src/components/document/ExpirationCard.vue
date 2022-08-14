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
        <h3 class="text-xl font-bold">
          <ion-icon name="calendar-outline" class="align-middle"></ion-icon>
          Valid until:
        </h3>
        <p class="text-xl text-right">{{ expiresAt }}</p>
      </div>
    </div>
  </div>
</template>
