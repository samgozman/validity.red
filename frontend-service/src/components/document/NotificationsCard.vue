<script setup lang="ts">
import { defineProps } from "vue";
import {
  NotificationsOutline,
  CheckmarkSharp,
  CloseSharp,
  AddOutline,
} from "@vicons/ionicons5";
import NotificationItem from "./NotificationItem.vue";
import type { INotification } from "./interfaces/INotification";
defineProps<{
  notifications: INotification[];
  documentId: string;
}>();
</script>

<template>
  <div class="card shadow-lg bg-base-100">
    <div class="card-body">
      <h2 class="my-4 text-xl font-bold card-title">
        <NotificationsOutline class="w-5 inline" />
        Notifications
      </h2>
      <p v-if="notifications.length === 0" class="text-center my-5">
        Add notifications..
      </p>
      <NotificationItem
        v-for:="notification in notifications"
        v-bind:key="notification.ID"
        v-bind:notification="notification"
        @refresh-notifications-event="refreshNotificationsEmit"
      />
      <form v-if="isFormActive" @submit.prevent="submit" class="mt-2">
        <div class="form-control">
          <div class="input-group">
            <input
              type="date"
              :min="minDate"
              class="input input-bordered w-[100vw]"
              v-model="inputDate"
            />
            <input
              type="time"
              class="input input-bordered w-[50vw]"
              v-model="inputTime"
            />
            <button class="btn btn-square btn-primary" type="submit">
              <CheckmarkSharp class="w-5" />
            </button>
            <button class="btn btn-square" @click="closeFromClicked">
              <CloseSharp class="w-5" />
            </button>
          </div>
        </div>
      </form>
      <div
        v-if="isAddButtonActive"
        class="justify-center space-x-2 card-actions"
      >
        <button class="btn btn-primary btn-circle" @click="addButtonClicked">
          <AddOutline class="w-5" />
        </button>
      </div>
      <div v-show="errorMsg" class="badge badge-error badge-outline w-full">
        {{ errorMsg }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { NotificationService } from "./NotificationService";
import { ErrorDecoder } from "@/services/ErrorDecoder";

export default defineComponent({
  data() {
    return {
      isFormActive: false,
      isAddButtonActive: true,
      minDate: new Date().toISOString().split("T")[0],
      inputDate: "",
      inputTime: "",
      errorMsg: "",
    };
  },
  methods: {
    refreshNotificationsEmit() {
      this.$emit("refreshNotificationsEvent");
    },
    async submit() {
      this.errorMsg = "";

      // TODO: Is timezone correct?
      const fullDate = new Date(`${this.inputDate}T${this.inputTime}:00`);
      if (isNaN(fullDate.getTime())) {
        this.errorMsg = "Invalid date.";
        return;
      }
      if (new Date(fullDate) < new Date()) {
        this.errorMsg = "Notification date is in the past!";
        return false;
      }
      try {
        await NotificationService.createOne({
          date: new Date(fullDate),
          documentId: this.documentId,
        });
      } catch (error) {
        this.errorMsg = await ErrorDecoder.decode(error);
      }

      this.closeFromClicked();
      this.$emit("refreshNotificationsEvent");
    },
    addButtonClicked() {
      this.isFormActive = true;
      this.isAddButtonActive = false;
    },
    closeFromClicked() {
      this.isFormActive = false;
      this.isAddButtonActive = true;
    },
  },
});
</script>
