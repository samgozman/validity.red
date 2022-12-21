<script setup lang="ts">
import { RouterLink } from "vue-router";
import ModalBig from "@/components/layout/ModalBig.vue";
defineProps<{
  modalId: string;
}>();
</script>

<template>
  <ModalBig :modalId="modalId">
    <p>Modal content</p>
  </ModalBig>
</template>
