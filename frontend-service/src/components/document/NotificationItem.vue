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
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

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
  },
  beforeMount() {
    this.setContext();
  },
});
</script>
