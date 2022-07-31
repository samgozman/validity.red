<script setup lang="ts">
import { defineProps } from "vue";
import type { INotification } from "./interfaces/INotification";
defineProps<{
  notification: INotification;
}>();
</script>

<template>
  <div class="indicator w-full">
    <div class="indicator-item indicator-middle indicator-end">
      <button class="btn btn-xs btn-accent btn-circle">X</button>
    </div>
    <div
      class="grid w-full h-10 rounded bg-primary text-primary-content place-content-center my-1"
    >
      {{ notification.date }}
    </div>
  </div>
</template>
