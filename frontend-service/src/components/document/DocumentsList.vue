<script setup lang="ts">
import DocumentListItem from "./DocumentListItem.vue";
</script>

<template>
  <div
    class="grid grid-cols-1 gap-6 lg:p-10 xl:grid-cols-3 lg:bg-base-200 rounded-box"
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
import axios from "axios";

import type { IDocument } from "./interfaces/IDocument";

interface DocumentGetAllResponse {
  error: boolean;
  message: string;
  data: {
    documents: IDocument[];
  };
}

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
      const payload = JSON.stringify({
        action: "DocumentGetAll",
      });

      try {
        const res = await axios.post<DocumentGetAllResponse>(
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

        const { error, message, data } = res.data;

        if (error) {
          this.error = true;
          this.errorMsg = message;
          return;
        }

        this.documents = data.documents;
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
