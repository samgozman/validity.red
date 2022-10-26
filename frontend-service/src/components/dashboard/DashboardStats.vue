<script setup lang="ts">
import UsedTypesItem from "@/components/dashboard/UsedTypesItem.vue";
import LatestDocumentsItem from "@/components/dashboard/LatestDocumentsItem.vue";
import CalendarMonth from "@/components/calendar/CalendarMonth.vue";
</script>

<template>
  <div
    v-if="calendarId"
    class="card col-span-full row-span-1 shadow-lg compact bg-base-100"
  >
    <div class="flex-col sm:flex-row items-center card-body">
      <p v-show="!errorMsgIcs">
        Export Validity.Red data to sync with your calendar app
      </p>
      <span v-show="errorMsgIcs" class="badge badge-error badge-outline w-full">
        Error: {{ errorMsgIcs }}
      </span>
      <button @click.prevent="getIcs" href="#" class="btn btn-primary btn-sm">
        Export calendar
      </button>
    </div>
  </div>
  <div
    class="card col-span-full row-span-1 shadow-lg compact bg-base-100 min-h-[40vh]"
  >
    <div class="flex-col sm:flex-row items-center card-body">
      <CalendarMonth />
    </div>
  </div>
  <div class="card col-span-3 row-span-3 shadow-lg compact bg-base-100">
    <div class="flex-col sm:flex-row items-center card-body">
      <div class="stat">
        <div class="stat-figure text-secondary stats-icon">
          <ion-icon name="documents-outline"></ion-icon>
        </div>
        <div class="stat-title">Documents</div>
        <div class="stat-value">{{ stats.totalDocuments }}</div>
        <div class="stat-desc">
          With {{ stats.usedTypes ? stats.usedTypes.length : 0 }} types
        </div>
      </div>

      <div class="divider divider-vertical sm:divider-horizontal"></div>

      <div class="stat">
        <div class="stat-figure text-secondary stats-icon">
          <ion-icon name="notifications-outline"></ion-icon>
        </div>
        <div class="stat-title">Notifications</div>
        <div class="stat-value">{{ stats.totalNotifications }}</div>
        <div class="stat-desc">
          Avg notifications per document - {{ avgNotifications.toFixed(2) }}
        </div>
      </div>
    </div>
  </div>
  <div class="card col-span-2 row-span-4 shadow-lg compact bg-base-100">
    <div class="flex-col items-center card-body">
      <h3 class="card-title text-primary text-left w-full px-4">
        Most used types
      </h3>
      <ul class="w-full sm:max-h-36 overflow-y-auto">
        <UsedTypesItem
          v-for="usedType in stats.usedTypes"
          v-bind:key="usedType.type"
          v-bind:usedType="usedType"
        />
      </ul>
    </div>
  </div>
  <div class="card col-span-2 row-span-4 shadow-lg compact bg-base-100">
    <div class="flex-col items-center card-body">
      <h3 class="card-title text-primary text-left w-full px-4">
        Latest documents
      </h3>
      <ul class="menu p-2 w-full sm:max-h-36 overflow-y-auto">
        <LatestDocumentsItem
          v-for="document in stats.latestDocuments"
          v-bind:key="document.ID"
          v-bind:document="document"
        />
      </ul>
    </div>
  </div>
  <div class="card col-span-3 row-span-1 shadow-lg compact bg-base-100">
    <div class="flex-col sm:flex-row items-center card-body">
      <p v-show="!errorMsg">App version and link to changelog</p>
      <span v-show="errorMsg" class="badge badge-error badge-outline w-full">
        Error: {{ errorMsg }}
      </span>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { DashboardService } from "./DashboardService";
import type { IDashboardStats } from "./interfaces/IDashboardStats";
import { ErrorDecoder } from "@/services/ErrorDecoder";

interface VueData {
  stats: IDashboardStats;
  avgNotifications: number;
  calendarId: string | null;
  errorMsg: string;
  errorMsgIcs: string;
}

export default defineComponent({
  data(): VueData {
    return {
      stats: {} as IDashboardStats,
      avgNotifications: 0,
      calendarId: localStorage.getItem("calendarId"),
      errorMsg: "",
      errorMsgIcs: "",
    };
  },
  methods: {
    async refresh() {
      this.errorMsg = "";
      try {
        this.stats = await DashboardService.getStats();
        this.avgNotifications =
          this.stats.totalNotifications / this.stats.totalDocuments || 0;
      } catch (error) {
        this.errorMsg = await ErrorDecoder.decode(error);
      }
    },
    async getIcs() {
      this.errorMsgIcs = "";
      try {
        const data = await DashboardService.getIcsFile(this.calendarId || "");
        // For some reason, the blob storage is not working properly
        // if we just set the href of existing vue-node to the blob url.
        // This is an old-fashioned workaround.
        const link = document.createElement("a");
        link.href = window.URL.createObjectURL(
          new Blob([data], { type: "text/calendar" })
        );
        link.setAttribute("download", "validity-calendar.ics"); //or any other extension
        document.body.appendChild(link);
        link.click();
        link.remove();
        // TODO: Fix! This way of downloading the file is not calling the download alert box.
        // And there for not trying to open the file in the calendar app. (and may not work on mobile)
        // So to fix this, we need to find a way to use a file server instead or proxy the original request link.
      } catch (error) {
        this.errorMsgIcs = await ErrorDecoder.decode(error);
      }
    },
  },
  beforeMount() {
    this.refresh();
  },
});
</script>

<style scoped>
.stats-icon {
  font-size: 2.25rem;
}
</style>
