<script setup lang="ts">
import { defineProps } from "vue";
import { RouterLink } from "vue-router";
import type { ICalendarNotification } from "./interfaces/ICalendarNotification";
defineProps<{
  notification: ICalendarNotification;
}>();
</script>

<template>
  <RouterLink
    :to="`/documents/${notification.documentID}`"
    class="flex items-center flex-shrink-0 h-3 md:h-5 md:px-1 text-ultra-small md:text-xs hover:bg-base-300"
  >
    <span class="font-light leading-none hidden md:block md:ml-2">
      {{
        new Date(notification.notificationDate).toLocaleTimeString("en-GB", {
          hour: "2-digit",
          minute: "2-digit",
        })
      }}
    </span>
    <span class="font-medium leading-none truncate md:ml-2">
      {{ notification.documentTitle }}
    </span>
  </RouterLink>
</template>

<style>
@media (max-width: 768px) {
  .text-ultra-small {
    font-size: 0.5rem;
    line-height: 0.75rem;
  }
}
</style>
