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
    :to="`/documents/${notification.documentId}`"
    class="flex items-center flex-shrink-0 h-5 px-1 text-xs hover:bg-base-300"
  >
    <span class="ml-2 font-light leading-none hidden md:block">
      {{
        new Date(notification.notificationDate).toLocaleTimeString("en-GB", {
          hour: "2-digit",
          minute: "2-digit",
        })
      }}
    </span>
    <span class="ml-2 font-medium leading-none truncate">
      {{ notification.documentTitle }}
    </span>
  </RouterLink>
</template>
