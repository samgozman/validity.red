<script setup lang="ts">
import { defineProps } from "vue";
import { RouterLink } from "vue-router";
import { DocumentType } from "@/components/document/DocumentType";
import type { IDocument } from "@/components/document/interfaces/IDocument";
defineProps<{
  document: IDocument;
}>();
</script>

<template>
  <li class="w-full">
    <RouterLink :to="`/documents/${document.ID}`" class="flex justify-between">
      <component
        :is="DocumentType.getIcon(document.type)"
        class="w-4"
      ></component>
      <span class="truncate">{{ document.title }}</span>
      <span>
        {{ new Date(document.expiresAt).toISOString().slice(0, 10) }}
      </span>
    </RouterLink>
  </li>
</template>
