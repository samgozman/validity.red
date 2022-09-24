<script setup lang="ts">
import CalendarDay from "./CalendarDay.vue";
</script>

<template>
  <div class="flex flex-grow w-full h-full md:h-screen xl:h-[60vh]">
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
        <h2 class="ml-2 text-xl font-bold leading-none">
          {{ currentDateString }}
        </h2>
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
        class="grid flex-grow w-full h-auto grid-cols-7 grid-rows-5 gap-px pt-px mt-1 bg-base-200 rounded-box"
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
import { CalendarService } from "./CalendarService";

export default defineComponent({
  data() {
    return {
      month: new Map<number, ICalendarNotification[]>(),
      currentDateString: "",
      error: false,
      errorMsg: "",
    };
  },
  methods: {
    async refresh() {
      try {
        const usersCalendar = await CalendarService.getCalendar();

        this.month = CalendarService.createCalendar(new Date(), usersCalendar);
        console.log(this.month);
      } catch (error) {
        this.error = true;
        this.errorMsg = "An error occurred, please try again";
        // TODO: Push error to Sentry
      }
    },
    setCurrentDateString(date: Date) {
      const month = date.toLocaleString("default", { month: "long" });
      const year = date.getFullYear();
      this.currentDateString = `${month}, ${year}`;
    },
  },
  beforeMount() {
    this.refresh();
    this.setCurrentDateString(new Date());
  },
});
</script>
