<script setup lang="ts">
import { defineProps } from "vue";
import ModalConfirmation from "../elements/ModalConfirmation.vue";
import type { IDocument } from "./interfaces/IDocument";
defineProps<{
  document: IDocument;
}>();
</script>

<template>
  <div class="card col-span-1 row-span-3 shadow-lg xl:col-span-2 bg-base-100">
    <div class="card-body">
      <h2 class="my-4 text-4xl font-bold card-title">{{ document.title }}</h2>
      <div class="mb-4 space-x-2 card-actions">
        <div class="badge badge-ghost">Type of document</div>
      </div>
      <p>
        {{ document.description }}
      </p>
      <div class="justify-end space-x-2 card-actions">
        <button class="btn btn-primary">Edit</button>
        <a :href="deleteAncor" class="btn">Delete</a>
      </div>
    </div>
    <ModalConfirmation
      :modalId="deleteModalId"
      message="Are you sure that you want to delete this document:"
      :actionName="document.title || ''"
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
      deleteAncor: `#delete-document`,
      deleteModalId: `delete-document`,
    };
  },
  methods: {
    async deleteDocument() {
      try {
        await DocumentService.deleteOne(this.document.ID);
        this.$router.push("/documents");
      } catch (error) {
        // TODO: push error to errors handler (display errors it in the UI)
        console.error("An error occurred, please try again", error);
        // TODO: Push error to Sentry
      }
    },
  },
});
</script>
