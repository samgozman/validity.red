<script setup lang="ts">
import DocumentCreate from "@/components/document/DocumentCreate.vue";
import StandardContainer from "@/components/layout/StandardContainer.vue";
</script>

<template>
  <StandardContainer>
    <DocumentCreate />
  </StandardContainer>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  created() {
    document.title = "Create document | Validity";
  },
});
</script>
