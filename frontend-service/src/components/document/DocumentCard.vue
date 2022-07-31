<script setup lang="ts">
import { defineProps } from "vue";
import type { IDocument } from "./interfaces/IDocument";
defineProps<{
  document: IDocument;
}>();
</script>

<template>
  <div class="card col-span-1 row-span-3 shadow-lg xl:col-span-2 bg-base-100">
    <div class="card-body">
      <h2 class="my-4 text-4xl font-bold card-title">{{ document.title }}</h2>
      <div class="mb-4 space-x-2 card-actions">
        <div class="badge badge-ghost">Type of document</div>
      </div>
      <p>
        {{ document.description }}
      </p>
      <div class="justify-end space-x-2 card-actions">
        <button class="btn btn-primary">Edit</button>
        <button class="btn">Delete</button>
      </div>
    </div>
  </div>
</template>
