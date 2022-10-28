<script setup lang="ts">
import { defineProps } from "vue";
import { RouterLink } from "vue-router";
import ModalConfirmation from "../elements/ModalConfirmation.vue";
import type { IDocument } from "./interfaces/IDocument";
import { DocumentType } from "./DocumentType";
defineProps<{
  document: IDocument;
}>();
</script>

<template>
  <div class="card shadow-lg bg-base-100">
    <div class="card-body">
      <h2 class="my-4 text-4xl font-bold card-title">{{ document.title }}</h2>
      <div class="mb-4 space-x-2 card-actions">
        <div class="badge badge-lg badge-ghost">
          <ion-icon
            :name="DocumentType.getIconStyle(document.type)"
            class="mr-1"
          ></ion-icon>
          {{ DocumentType.getName(document.type) }}
        </div>
      </div>
      <p>
        {{ document.description }}
      </p>
      <div class="justify-end space-x-2 card-actions">
        <RouterLink
          :to="`/documents/${document.ID}/edit`"
          class="btn btn-primary"
          :v-model="document.ID"
          >Edit</RouterLink
        >
        <a :href="deleteAnchor" class="btn">Delete</a>
      </div>
      <div v-show="errorMsg" class="badge badge-error badge-outline w-full">
        {{ errorMsg }}
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
import { ErrorDecoder } from "@/services/ErrorDecoder";

export default defineComponent({
  data() {
    return {
      deleteAnchor: "#delete-document",
      deleteModalId: "delete-document",
      errorMsg: "",
    };
  },
  methods: {
    async deleteDocument() {
      this.errorMsg = "";
      try {
        await DocumentService.deleteOne(this.document.ID);
        this.$router.push("/documents");
      } catch (error) {
        this.errorMsg = await ErrorDecoder.decode(error);
      }
    },
  },
});
</script>
