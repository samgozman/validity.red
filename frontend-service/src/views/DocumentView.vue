<script setup lang="ts">
import DocumentFullBlock from "@/components/document/DocumentFullBlock.vue";
import StandardContainer from "@/components/layout/StandardContainer.vue";
</script>

<template>
  <StandardContainer>
    <DocumentFullBlock @setDocumentTitle="setDocumentTitle" />
  </StandardContainer>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  methods: {
    setDocumentTitle(title: string) {
      document.title = title + " | Validity";
    },
  },
  created() {
    document.title = "Document | Validity";
  },
});
</script>
