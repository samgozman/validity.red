<script setup lang="ts">
import { defineProps } from "vue";
import { DocumentType } from "@/components/document/DocumentType";
import type { IDocument } from "@/components/document/interfaces/IDocument";
defineProps<{
  document: IDocument;
}>();
</script>

<template>
  <li>
    <a>
      <ion-icon :name="DocumentType.getIconStyle(document.type)"></ion-icon>
      {{ document.title }}
      <span>
        {{ new Date(document.expiresAt).toISOString().slice(0, 10) }}
      </span>
    </a>
  </li>
</template>
