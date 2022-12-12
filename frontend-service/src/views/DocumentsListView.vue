<script setup lang="ts">
import DocumentsList from "@/components/document/DocumentsList.vue";
import StandardContainer from "@/components/layout/StandardContainer.vue";
</script>

<template>
  <StandardContainer>
    <DocumentsList />
  </StandardContainer>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  created() {
    document.title = "Documents | Validity.Red";
  },
});
</script>
