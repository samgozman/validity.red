<script setup lang="ts">
import { defineProps } from "vue";
defineProps<{
  msg: string;
}>();
</script>

<template>
  <button class="btn">{{ msg }}</button>
</template>
