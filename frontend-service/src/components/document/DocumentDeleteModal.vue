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
        <a href="#" class="btn btn-primary" @click.prevent="clickConfirm"
          >Confirm</a
        >
        <a href="#" class="btn">Close</a>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  methods: {
    async clickConfirm() {
      this.$emit("deleteDocumentEvent");
    },
  },
});
</script>
