<script setup lang="ts">
import CalendarDay from "./CalendarDay.vue";
</script>

<template>
  <div class="flex flex-grow w-full h-screen overflow-auto">
    <div class="flex flex-col flex-grow">
      <!-- Header -->
      <div class="flex items-center mt-4">
        <div class="flex ml-6">
          <button>
            <ion-icon name="chevron-back-outline"></ion-icon>
          </button>
          <button>
            <ion-icon name="chevron-forward-outline"></ion-icon>
          </button>
        </div>
        <h2 class="ml-2 text-xl font-bold leading-none">September, 2020</h2>
      </div>
      <!-- Day of the week columns -->
      <div class="grid grid-cols-7 mt-4">
        <div class="pl-1 text-sm">Mon</div>
        <div class="pl-1 text-sm">Tue</div>
        <div class="pl-1 text-sm">Wed</div>
        <div class="pl-1 text-sm">Thu</div>
        <div class="pl-1 text-sm">Fri</div>
        <div class="pl-1 text-sm">Sat</div>
        <div class="pl-1 text-sm">Sun</div>
      </div>
      <!-- Main block -->
      <div
        class="grid flex-grow w-full h-auto grid-cols-7 grid-rows-5 gap-px pt-px mt-1 bg-base-200"
      >
        <div>
          <!-- Placeholder -->
        </div>
        <CalendarDay
          v-for="day in month"
          v-bind:key="day[0]"
          v-bind:indexDay="day[0]"
          v-bind:notifications="day[1]"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import type { ICalendarNotification } from "./interfaces/ICalendarNotification";

const notifications1: ICalendarNotification[] = [
  {
    documentId: "8fd1160f-60b7-4396-8a14-ad0eabaf8ee9",
    documentTitle: "Document placeholder",
    notificationDate: new Date().toISOString(),
  },
  {
    documentId: "8fd1160f-60b7-4396-8a14-ad0eabaf8ee9",
    documentTitle: "Document placeholder",
    notificationDate: new Date().toISOString(),
  },
];

const notifications2: ICalendarNotification[] = [
  {
    documentId: "28292f09-0a55-4ecd-9e4d-eb234f4b23ed",
    documentTitle: "Expiration of visa",
    notificationDate: new Date().toISOString(),
  },
];

// Placeholder month
const month = new Map<number, ICalendarNotification[]>([
  [1, []],
  [2, []],
  [3, []],
  [4, []],
  [5, []],
  [6, []],
  [7, notifications1],
  [8, []],
  [9, []],
  [10, []],
  [11, []],
  [12, []],
  [13, []],
  [14, []],
  [15, []],
  [16, []],
  [17, []],
  [18, []],
  [19, []],
  [20, []],
  [21, []],
  [22, []],
  [23, []],
  [24, []],
  [25, notifications2],
  [26, []],
  [27, []],
  [28, []],
  [29, []],
  [30, []],
  [31, []],
]);

export default defineComponent({
  data() {
    return {
      month,
    };
  },
  methods: {},
});
</script>
