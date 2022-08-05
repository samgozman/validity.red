<script setup lang="ts">
import { defineProps } from "vue";
import { RouterLink } from "vue-router";
import type { IDocument } from "./interfaces/IDocument";
import ModalConfirmation from "../elements/ModalConfirmation.vue";
defineProps<{
  document: IDocument;
}>();
</script>

<template>
  <div class="card shadow-lg compact side bg-base-100">
    <div class="flex-row items-center space-x-4 card-body">
      <div class="flex-1">
        <h2 class="card-title text-primary">{{ document.title }}</h2>
        <p class="text-base-content text-opacity-80">
          {{ document.description }}
        </p>
        <p class="text-base-content text-opacity-40">01.01.1970</p>
      </div>
      <div class="flex space-x-2 flex-0">
        <RouterLink
          :to="documentLink"
          class="btn btn-sm btn-square btn-primary"
        >
          v
        </RouterLink>
        <!-- TODO: call delete method after pop-up confirmation -->
        <a :href="deleteAncor" class="btn btn-sm btn-square">d</a>
      </div>
    </div>
    <ModalConfirmation
      :modalId="deleteModalId"
      message="Are you sure that you want to delete this document:"
      :actionName="document.title"
      @confirmEvent="deleteDocument"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { DocumentService } from "./DocumentService";

export default defineComponent({
  data() {
    return {
      deleteAncor: "",
      deleteModalId: "",
      documentLink: "",
    };
  },
  methods: {
    setContext() {
      this.deleteAncor = `#delete-${this.document.ID}`;
      this.deleteModalId = `delete-${this.document.ID}`;
      this.documentLink = `documents/${this.document.ID}`;
    },
    async deleteDocument() {
      try {
        await DocumentService.deleteOne(this.document.ID);
        this.$router.push("/documents");
        this.$emit("refreshDocumentsEvent");
      } catch (error) {
        // TODO: push error to errors handler (display errors it in the UI)
        console.error("An error occurred, please try again", error);
        // TODO: Push error to Sentry
      }
    },
  },
  beforeMount() {
    this.setContext();
  },
});
</script>
