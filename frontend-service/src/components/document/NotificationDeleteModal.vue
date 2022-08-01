deleteOne
<script setup lang="ts">
import { defineProps } from "vue";
import type { INotification } from "./interfaces/INotification";
defineProps<{
  modalId: string;
  notification: INotification;
}>();
</script>

<template>
  <div v-bind:id="modalId" class="modal">
    <div class="modal-box">
      <p>
        Are you sure that you want to delete this notification:
        <strong>{{ notification.date }}</strong> ?
      </p>
      <div class="modal-action">
        <a href="#" class="btn btn-primary" @click.prevent="deleteNotification"
          >Confirm</a
        >
        <a href="#" class="btn">Close</a>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { NotificationService } from "./NotificationService";

export default defineComponent({
  methods: {
    async deleteNotification() {
      try {
        await NotificationService.deleteOne({
          id: this.notification.ID,
          documentId: this.notification.documentID,
        });
        // TODO: Close modal and update notifications list
      } catch (error) {
        // TODO: push error to errors handler (display errors it in the UI)
        console.error("An error occurred, please try again", error);
        // TODO: Push error to Sentry
      }
    },
  },
});
</script>
