<script setup lang="ts">
import { defineProps } from "vue";
import NotificationItem from "./NotificationItem.vue";
import type { INotification } from "./interfaces/INotification";
defineProps<{
  notifications: INotification[];
}>();
</script>

<template>
  <div class="card col-span-1 row-span-3 shadow-lg bg-base-100">
    <div class="card-body">
      <h2 class="my-4 text-xl font-bold card-title">Notifications</h2>
      <p v-if="notifications.length === 0" class="text-center my-5">
        Add notifications..
      </p>
      <div>
        <NotificationItem
          v-for="notification in notifications"
          v-bind:key="notification.ID"
          v-bind:notification="notification"
        />
      </div>
      <div class="justify-center space-x-2 card-actions">
        <button class="btn btn-primary btn-circle">+</button>
      </div>
    </div>
  </div>
</template>
