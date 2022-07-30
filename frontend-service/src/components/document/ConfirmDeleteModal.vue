deleteOne
<script setup lang="ts">
import { defineProps } from "vue";
import type { IDocument } from "./interfaces/IDocument";
defineProps<{
  modalId: string;
  document: IDocument;
}>();
</script>

<template>
  <div v-bind:id="modalId" class="modal">
    <div class="modal-box">
      <p>
        Are you sure that you want to delete this document:
        <strong>{{ document.title }}</strong> ?
      </p>
      <div class="modal-action">
        <a href="#" class="btn btn-primary" @click.prevent="deleteDocument"
          >Confirm</a
        >
        <a href="#" class="btn">Close</a>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { DocumentService } from "./DocumentService";

export default defineComponent({
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
