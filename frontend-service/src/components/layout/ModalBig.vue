<script setup lang="ts">
import { defineProps } from "vue";
defineProps<{
  modalId: string;
}>();
</script>

<template>
  <div v-bind:id="modalId" class="modal">
    <div class="modal-box">
      <slot />
      <div class="modal-action justify-center">
        <a href="#" class="btn btn-primary w-full">Ok 👌</a>
      </div>
    </div>
  </div>
</template>
