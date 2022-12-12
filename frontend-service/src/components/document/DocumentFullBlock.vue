<script setup lang="ts">
import DocumentCard from "./DocumentCard.vue";
import NotificationsCard from "./NotificationsCard.vue";
import DocumentBreadcrumbs from "./DocumentBreadcrumbs.vue";
import ExpirationCard from "./ExpirationCard.vue";
</script>

<template>
  <DocumentBreadcrumbs :title="document.title || ''" class="lg:px-10 pt-7" />
  <div v-show="errorMsg" class="badge badge-error badge-outline w-full">
    Error: {{ errorMsg }}
  </div>
  <div
    class="grid grid-cols-1 gap-6 py-5 lg:px-10 xl:grid-cols-3 lg:bg-base-200 rounded-box"
  >
    <DocumentCard
      :document="document"
      class="card col-span-1 row-span-3 xl:col-span-2"
    />
    <ExpirationCard
      :expiresAt="document.expiresAt || ''"
      class="card col-span-1 row-span-1"
    />
    <NotificationsCard
      class="card col-span-1 row-span-2"
      :notifications="notifications"
      :documentId="documentId"
      @refresh-notifications-event="fetchNotifications"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { DocumentService } from "./DocumentService";
import { NotificationService } from "./NotificationService";
import type { IDocument } from "./interfaces/IDocument";
import type { INotification } from "./interfaces/INotification";
import { ErrorDecoder } from "@/services/ErrorDecoder";

interface VueData {
  document: IDocument;
  notifications: INotification[];
  documentId: string;
  errorMsg: string;
}

export default defineComponent({
  data(): VueData {
    return {
      document: {} as IDocument,
      notifications: [],
      documentId:
        typeof this.$route.params.id === "string" ? this.$route.params.id : "",
      errorMsg: "",
    };
  },
  methods: {
    async fetchDocument() {
      try {
        this.errorMsg = "";
        this.document = await DocumentService.getOne(this.documentId);
        this.$emit("setDocumentTitle", this.document.title);
      } catch (error) {
        this.errorMsg = await ErrorDecoder.decode(error, this.$router);
      }
    },
    async fetchNotifications() {
      try {
        this.notifications = await NotificationService.getAll(this.documentId);
      } catch (error) {
        this.errorMsg = await ErrorDecoder.decode(error);
      }
    },
  },
  async beforeMount() {
    await this.fetchDocument();
    await this.fetchNotifications();
  },
});
</script>
