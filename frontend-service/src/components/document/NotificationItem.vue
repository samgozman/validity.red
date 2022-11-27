<script setup lang="ts">
import { defineProps } from "vue";
import { CloseOutline } from "@vicons/ionicons5";
import ModalConfirmation from "../elements/ModalConfirmation.vue";
import type { INotification } from "./interfaces/INotification";
defineProps<{
  notification: INotification;
}>();
</script>

<template>
  <div class="indicator w-full">
    <div class="indicator-item indicator-middle indicator-end">
      <a :href="deleteAnchor" class="btn btn-xs btn-accent btn-circle">
        <CloseOutline class="w-5" />
      </a>
    </div>
    <div
      class="grid w-full h-10 rounded bg-primary text-primary-content place-content-center my-1"
    >
      <p v-if="!errorMsg">{{ dateWithTZ }}</p>
      <p v-if="errorMsg">Error: {{ errorMsg }}</p>
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
import { ErrorDecoder } from "@/services/ErrorDecoder";

export default defineComponent({
  data() {
    return {
      dateWithTZ: new Date(this.notification.date).toLocaleString("en-GB", {
        timeZoneName: "short",
      }),
      deleteAnchor: "",
      deleteModalId: "",
      errorMsg: "",
    };
  },
  methods: {
    setContext() {
      this.deleteAnchor = `#delete-${this.notification.ID}`;
      this.deleteModalId = `delete-${this.notification.ID}`;
    },
    async deleteNotification() {
      this.errorMsg = "";
      try {
        await NotificationService.deleteOne({
          id: this.notification.ID,
          documentId: this.notification.documentID,
        });
        // push?
        this.$emit("refreshNotificationsEvent");
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
