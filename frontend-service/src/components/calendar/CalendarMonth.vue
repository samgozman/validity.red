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
        <CalendarDay v-for="day in 30" v-bind:key="day" v-bind:day="day" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  data() {
    return {};
  },
  methods: {},
});
</script>
