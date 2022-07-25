<script setup lang="ts">
import { defineProps } from "vue";
defineProps<{
  label: string;
}>();
</script>

<template>
  <label class="label">
    <span class="label-text">{{ label }}</span>
  </label>
</template>
