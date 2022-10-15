<script setup lang="ts">
import CalendarDay from "./CalendarDay.vue";
</script>

<template>
  <div class="flex flex-grow w-full h-[60vh] md:h-[70vh] xl:h-[60vh]">
    <div class="flex flex-col flex-grow">
      <!-- Header -->
      <div class="flex items-center mt-4">
        <div class="flex ml-6">
          <button @click.prevent="prevMonth">
            <ion-icon name="chevron-back-outline"></ion-icon>
          </button>
          <button @click.prevent="nextMonth">
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
        class="grid flex-grow w-full h-auto grid-cols-7 gap-px pt-px mt-1 bg-base-200 rounded-box"
      >
        <!-- Placeholder -->
        <div v-for="day in currentFirstDayOfWeek" v-bind:key="day" />
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
      currentFirstDayOfWeek: 0,
      currentDate: new Date(),
      currentDateString: "",
      error: false,
      errorMsg: "",
    };
  },
  methods: {
    async refresh() {
      this.setCurrentDateString(this.currentDate);

      try {
        const usersCalendar = await CalendarService.getCalendar();

        this.month = CalendarService.createCalendar(
          this.currentDate,
          usersCalendar
        );
      } catch (error) {
        this.error = true;
        this.errorMsg = "An error occurred, please try again";
        // TODO: Push error to Sentry
      }
    },
    prevMonth() {
      this.currentDate.setMonth(this.currentDate.getMonth() - 1);
      this.refresh();
    },
    nextMonth() {
      this.currentDate.setMonth(this.currentDate.getMonth() + 1);
      this.refresh();
    },
    setCurrentDateString(date: Date) {
      const month = date.toLocaleString("default", { month: "long" });
      const year = date.getFullYear();

      this.currentDate = date;
      this.currentDateString = `${month}, ${year}`;

      // To convert getDay from 0-6 (Sun - Sat) to 0-6 (Mon - Sun)
      const week = new Map<number, number>([
        [0, 6],
        [1, 0],
        [2, 1],
        [3, 2],
        [4, 3],
        [5, 4],
        [6, 5],
      ]);
      const dayOfWeek = new Date(
        date.getFullYear(),
        date.getMonth(),
        1
      ).getDay();

      this.currentFirstDayOfWeek = week.get(dayOfWeek) || 0;
    },
  },
  beforeMount() {
    this.refresh();
  },
});
</script>
