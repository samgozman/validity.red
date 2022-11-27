<script setup lang="ts">
import { defineProps } from "vue";
import { DocumentType } from "@/components/document/DocumentType";
import type { IUsedType } from "./interfaces/IUsedType";
defineProps<{
  usedType: IUsedType;
}>();
</script>

<template>
  <li class="px-4 py-1">
    <a class="flex justify-between">
      <span>
        <component
          :is="DocumentType.getIcon(usedType.type)"
          class="w-4 inline"
        ></component>
        {{ DocumentType.getName(usedType.type) }}
      </span>
      <span>{{ usedType.count }}</span>
    </a>
  </li>
</template>
