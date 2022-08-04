<script setup lang="ts">
import { defineProps } from "vue";
import NotificationDeleteModal from "./NotificationDeleteModal.vue";
import type { INotification } from "./interfaces/INotification";
defineProps<{
  notification: INotification;
}>();
</script>

<template>
  <div class="indicator w-full">
    <div class="indicator-item indicator-middle indicator-end">
      <a :href="deleteAncor" class="btn btn-xs btn-accent btn-circle">X</a>
    </div>
    <div
      class="grid w-full h-10 rounded bg-primary text-primary-content place-content-center my-1"
    >
      {{ notification.date }}
    </div>
    <NotificationDeleteModal
      :modalId="deleteModalId"
      :notification="notification"
      @delete-notification-event="deleteNotification"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { NotificationService } from "./NotificationService";

export default defineComponent({
  data() {
    return {
      deleteAncor: "",
      deleteModalId: "",
    };
  },
  methods: {
    setContext() {
      this.deleteAncor = `#delete-${this.notification.ID}`;
      this.deleteModalId = `delete-${this.notification.ID}`;
    },
    async deleteNotification() {
      try {
        await NotificationService.deleteOne({
          id: this.notification.ID,
          documentId: this.notification.documentID,
        });
        // push?
        this.$emit("refreshNotificationsEvent");
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
