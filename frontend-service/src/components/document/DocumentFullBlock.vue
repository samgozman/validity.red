<script setup lang="ts">
import DocumentCard from "./DocumentCard.vue";
import NotificationsCard from "./NotificationsCard.vue";
import DocumentBreadcrumbs from "./DocumentBreadcrumbs.vue";
</script>

<template>
  <DocumentBreadcrumbs :title="document.title || ''" class="lg:px-10 pt-7" />
  <div
    class="grid grid-cols-1 gap-6 py-5 lg:px-10 xl:grid-cols-3 lg:bg-base-200 rounded-box"
  >
    <DocumentCard :document="document" />
    <NotificationsCard :notifications="notifications" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { DocumentService } from "./DocumentService";
import type { IDocument } from "./interfaces/IDocument";
import type { INotification } from "./interfaces/INotification";

interface VueData {
  document: IDocument;
  notifications: INotification[];
}

export default defineComponent({
  data(): VueData {
    return {
      document: {} as IDocument,
      notifications: [],
    };
  },
  methods: {
    async refresh() {
      try {
        const docId: string =
          typeof this.$route.params.id === "string"
            ? this.$route.params.id
            : "";

        const doc = await DocumentService.getOne(docId);
        this.document = doc.document;
        this.notifications = doc.notifications || [];
      } catch (error) {
        // TODO: Push error to Sentry
        // TODO: Navigate to 404 page
      }
    },
  },
  beforeMount() {
    this.refresh();
  },
});
</script>
