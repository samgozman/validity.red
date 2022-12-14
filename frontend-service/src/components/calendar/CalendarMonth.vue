<script setup lang="ts">
import CalendarDay from "./CalendarDay.vue";
import { ChevronBackOutline, ChevronForwardOutline } from "@vicons/ionicons5";
</script>

<template>
  <div class="flex flex-grow w-full">
    <div class="flex flex-col flex-grow">
      <!-- Header -->
      <div class="flex items-center mt-4">
        <div class="flex">
          <button @click.prevent="prevMonth">
            <ChevronBackOutline class="w-6" />
          </button>
          <button @click.prevent="nextMonth">
            <ChevronForwardOutline class="w-6" />
          </button>
        </div>
        <h2 v-show="!errorMsg" class="ml-2 text-xl font-bold leading-none">
          {{ currentDateString }}
        </h2>
        <h2 v-show="errorMsg" class="badge badge-error badge-outline w-full">
          Error: {{ errorMsg }}
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
import { ErrorDecoder } from "@/services/ErrorDecoder";
import { setCalendarCurrentDate, state } from "@/state";

export default defineComponent({
  data() {
    return {
      month: new Map<number, ICalendarNotification[]>(),
      currentFirstDayOfWeek: 0,
      currentDate: new Date(),
      currentDateString: "",
      errorMsg: "",
    };
  },
  methods: {
    async refresh() {
      this.currentDate = state.value.user.calendarCurrentDate
        ? new Date(state.value.user.calendarCurrentDate)
        : new Date();
      setCalendarCurrentDate(this.currentDate.toString());
      this.setCurrentDateString(this.currentDate);
      this.errorMsg = "";
      try {
        const usersCalendar = await CalendarService.getCalendar();

        this.month = CalendarService.createCalendar(
          this.currentDate,
          usersCalendar
        );
      } catch (error) {
        this.errorMsg = await ErrorDecoder.decode(error);
      }
    },
    prevMonth() {
      this.currentDate.setMonth(this.currentDate.getMonth() - 1);
      setCalendarCurrentDate(this.currentDate.toString());
      this.refresh();
    },
    nextMonth() {
      this.currentDate.setMonth(this.currentDate.getMonth() + 1);
      setCalendarCurrentDate(this.currentDate.toString());
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
