<script setup lang="ts">
import DocumentListItem from "./DocumentListItem.vue";
</script>

<template>
  <div
    class="grid grid-cols-1 gap-6 py-6 lg:p-10 xl:grid-cols-3 lg:bg-base-200 rounded-box"
  >
    <DocumentListItem
      v-for="document in documents"
      v-bind:key="document.ID"
      v-bind:document="document"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { DocumentService } from "./DocumentService";
import type { IDocument } from "./interfaces/IDocument";

interface VueData {
  documents: IDocument[];
  error: boolean;
  errorMsg: string;
}

export default defineComponent({
  data(): VueData {
    return {
      documents: [],
      error: false,
      errorMsg: "",
    };
  },
  methods: {
    async refresh() {
      try {
        const documents = await DocumentService.getAll();
        this.documents = documents;
      } catch (error) {
        this.error = true;
        this.errorMsg = "An error occurred, please try again";
        // TODO: Push error to Sentry
      }
    },
  },
  beforeMount() {
    this.refresh();
  },
});
</script>
