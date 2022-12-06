<script setup lang="ts">
import DocumentListItem from "./DocumentListItem.vue";
import { RouterLink } from "vue-router";
import { AddOutline } from "@vicons/ionicons5";
</script>

<template>
  <div v-show="errorMsg" class="badge badge-error badge-outline w-full">
    Error: {{ errorMsg }}
  </div>
  <div
    v-show="isDocumentsAvailable"
    class="grid grid-cols-1 gap-6 py-6 lg:p-10 md:grid-cols-2 xl:grid-cols-3 lg:bg-base-200 rounded-box"
  >
    <DocumentListItem
      v-for="document in documents"
      v-bind:key="document.ID"
      v-bind:document="document"
      @refresh-documents-event="refresh"
    />
  </div>
  <div
    v-show="!isDocumentsAvailable"
    class="grid items-center justify-items-center h-full"
  >
    <RouterLink
      class="btn btn-lg mr-4 rounded-full w-max px-4"
      to="/documents/create"
    >
      <AddOutline class="w-6 lg:mr-1" />
      <span class="lg:block">Add first document</span>
    </RouterLink>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { DocumentService } from "./DocumentService";
import type { IDocument } from "./interfaces/IDocument";
import { ErrorDecoder } from "@/services/ErrorDecoder";

interface VueData {
  documents: IDocument[];
  isDocumentsAvailable: boolean;
  errorMsg: string;
}

export default defineComponent({
  data(): VueData {
    return {
      documents: [],
      // To fix button from flickering
      isDocumentsAvailable: true,
      errorMsg: "",
    };
  },
  methods: {
    async refresh() {
      this.errorMsg = "";
      try {
        const documents = await DocumentService.getAll();
        this.documents = documents;

        this.isDocumentsAvailable = documents.length > 0;
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
