<script setup lang="ts">
import DocumentListItem from "./DocumentListItem.vue";
</script>

<template>
  <div v-show="errorMsg" class="badge badge-error badge-outline w-full">
    Error: {{ errorMsg }}
  </div>
  <div
    class="grid grid-cols-1 gap-6 py-6 lg:p-10 md:grid-cols-2 xl:grid-cols-3 lg:bg-base-200 rounded-box"
  >
    <DocumentListItem
      v-for="document in documents"
      v-bind:key="document.ID"
      v-bind:document="document"
      @refresh-documents-event="refresh"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { DocumentService } from "./DocumentService";
import type { IDocument } from "./interfaces/IDocument";
import { ErrorDecoder } from "@/services/ErrorDecoder";

interface VueData {
  documents: IDocument[];
  errorMsg: string;
}

export default defineComponent({
  data(): VueData {
    return {
      documents: [],
      errorMsg: "",
    };
  },
  methods: {
    async refresh() {
      this.errorMsg = "";
      try {
        const documents = await DocumentService.getAll();
        this.documents = documents;
      } catch (error) {
        this.errorMsg = await ErrorDecoder.decode(error);
      }
    },
  },
  beforeMount() {
    this.refresh();
  },
});
</script>
