<script setup lang="ts">
import { defineProps } from "vue";
import type { ICalendarNotification } from "./interfaces/ICalendarNotification";
import CalendarNotificationItem from "./CalendarNotificationItem.vue";
defineProps<{
  /** Day of the month number */
  indexDay: number;
  notifications: ICalendarNotification[];
}>();
</script>

<template>
  <div class="relative flex flex-col bg-base-100 group h-12 md:h-24">
    <span class="mx-1 md:mx-2 my-1 text-xs font-bold">{{ indexDay }}</span>
    <!-- Notifications -->
    <div class="flex flex-col px-1 md:py-1 overflow-auto">
      <CalendarNotificationItem
        v-for="(item, index) in notifications"
        v-bind:key="index"
        v-bind:notification="item"
      />
    </div>
  </div>
</template>
