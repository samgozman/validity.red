<script setup lang="ts">
import { defineProps } from "vue";
defineProps<{
  day: number;
}>();
</script>

<template>
  <div class="relative flex flex-col bg-base-100 group">
    <span class="mx-2 my-1 text-xs font-bold">{{ day }}</span>
    <!-- Notifications -->
    <div class="flex flex-col px-1 py-1 overflow-auto">
      <button
        class="flex items-center flex-shrink-0 h-5 px-1 text-xs hover:bg-base-300"
      >
        <span class="ml-2 font-light leading-none">8:30am</span>
        <span class="ml-2 font-medium leading-none truncate"
          >An unconfirmed event</span
        >
      </button>
      <button
        class="flex items-center flex-shrink-0 h-5 px-1 text-xs hover:bg-base-300"
      >
        <span class="ml-2 font-light leading-none">2:15pm</span>
        <span class="ml-2 font-medium leading-none truncate">
          A confirmed event
        </span>
      </button>
    </div>
  </div>
</template>
