<script setup lang="ts">
import { defineProps } from "vue";
import { RouterLink } from "vue-router";
import { TrashOutline } from "@vicons/ionicons5";
import type { IDocument } from "./interfaces/IDocument";
import ModalConfirmation from "../elements/ModalConfirmation.vue";
import { DocumentType } from "./DocumentType";
defineProps<{
  document: IDocument;
}>();
</script>

<template>
  <div
    class="card shadow-lg compact bg-base-100"
    @mouseover="isHovering = true"
    @mouseout="isHovering = false"
  >
    <div class="flex-row items-center space-x-4 card-body">
      <div class="flex-1 w-[80%]">
        <RouterLink :to="documentLink" class="card-title text-primary">
          <component
            :is="DocumentType.getIcon(document.type)"
            class="w-5 text-base-content"
          ></component>
          {{ document.title }}
        </RouterLink>
        <p class="text-base-content text-opacity-80 truncate">
          {{ document.description }}
        </p>
        <p v-show="errorMsg" class="badge badge-error badge-outline w-full">
          Error: {{ errorMsg }}
        </p>
        <p class="text-base-content text-opacity-40">
          {{ document.expiresAt }}
        </p>
      </div>
      <div class="flex space-x-2 flex-0" :class="{ 'opacity-0': !isHovering }">
        <a :href="deleteAnchor" class="btn btn-sm btn-circle">
          <TrashOutline class="w-5" />
        </a>
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
import { ErrorDecoder } from "@/services/ErrorDecoder";

export default defineComponent({
  data() {
    return {
      isHovering: false,
      deleteAnchor: "",
      deleteModalId: "",
      documentLink: "",
      errorMsg: "",
    };
  },
  methods: {
    setContext() {
      this.deleteAnchor = `#delete-${this.document.ID}`;
      this.deleteModalId = `delete-${this.document.ID}`;
      this.documentLink = `documents/${this.document.ID}`;
    },
    async deleteDocument() {
      this.errorMsg = "";
      try {
        await DocumentService.deleteOne(this.document.ID);
        this.$router.push("/documents");
        this.$emit("refreshDocumentsEvent");
      } catch (error) {
        this.errorMsg = await ErrorDecoder.decode(error);
      }
    },
  },
  beforeMount() {
    this.setContext();
  },
});
</script>
