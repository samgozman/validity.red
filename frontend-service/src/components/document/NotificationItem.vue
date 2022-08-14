<script setup lang="ts">
import { defineProps } from "vue";
import ModalConfirmation from "../elements/ModalConfirmation.vue";
import type { INotification } from "./interfaces/INotification";
defineProps<{
  notification: INotification;
}>();
</script>

<template>
  <div class="indicator w-full">
    <div class="indicator-item indicator-middle indicator-end">
      <a :href="deleteAncor" class="btn btn-xs btn-accent btn-circle">
        <ion-icon name="close-outline" class="text-lg"></ion-icon>
      </a>
    </div>
    <div
      class="grid w-full h-10 rounded bg-primary text-primary-content place-content-center my-1"
    >
      {{ dateWithTZ }}
    </div>
    <ModalConfirmation
      :modalId="deleteModalId"
      message="Are you sure that you want to delete this notification:"
      :actionName="dateWithTZ"
      @confirmEvent="deleteNotification"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { NotificationService } from "./NotificationService";

export default defineComponent({
  data() {
    return {
      dateWithTZ: new Date(this.notification.date).toLocaleString("en-GB", {
        timeZoneName: "short",
      }),
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
