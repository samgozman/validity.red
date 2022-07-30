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
        <a href="#" class="btn btn-primary" @click="deleteDocument">Confirm</a>
        <a href="#" class="btn">Close</a>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import axios from "axios";

interface DocumentDeleteResponse {
  error: boolean;
  message: string;
}

export default defineComponent({
  methods: {
    async deleteDocument(e: Event) {
      e.preventDefault();

      const payload = JSON.stringify({
        action: "DocumentDelete",
        document: {
          id: this.document.ID,
        },
      });

      try {
        const res = await axios.post<DocumentDeleteResponse>(
          `http://localhost:8080/handle`,
          payload,
          {
            // To pass Set-Cookie header
            withCredentials: true,
            // Handle 401 error like a normal situation
            validateStatus: (status) =>
              (status >= 200 && status < 300) || status === 401,
          }
        );

        const { error, message } = res.data;

        if (error) {
          throw new Error(message);
        }
      } catch (error) {
        // TODO: push error to errors handler (display errors it in the UI)
        console.error("An error occurred, please try again", error);
        // TODO: Push error to Sentry
      }
    },
  },
});
</script>
