<script setup lang="ts">
import { defineProps } from "vue";
defineProps<{
  modalId: string;
  documentTitle: string;
}>();
</script>

<template>
  <div v-bind:id="modalId" class="modal">
    <div class="modal-box">
      <p>
        Are you sure that you want to delete this document:
        <strong>{{ documentTitle }}</strong> ?
      </p>
      <div class="modal-action">
        <a href="#" class="btn btn-primary">Confirm</a>
        <a href="#" class="btn">Close</a>
      </div>
    </div>
  </div>
</template>
