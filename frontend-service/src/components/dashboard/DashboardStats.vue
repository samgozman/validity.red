<script setup lang="ts">
import { RouterLink } from "vue-router";
import { DocumentOutline, NotificationsOutline } from "@vicons/ionicons5";
import UsedTypesItem from "@/components/dashboard/UsedTypesItem.vue";
import LatestDocumentsItem from "@/components/dashboard/LatestDocumentsItem.vue";
import CalendarMonth from "@/components/calendar/CalendarMonth.vue";
</script>

<template>
  <div
    class="hidden xl:flex card xl:col-span-4 row-span-1 shadow-lg compact bg-base-100"
  >
    <div class="flex-row items-center card-body">
      <h3 class="text-xl font-bold">Sync with your calendar app</h3>
    </div>
  </div>
  <div
    class="card col-span-full xl:col-span-3 row-span-1 shadow-lg compact bg-base-100"
  >
    <div class="flex-col items-center card-body">
      <h3 class="xl:hidden text-center sm:text-left text-xl font-bold w-full">
        Sync with your calendar app
      </h3>
      <div class="form-control w-full md:flex-row">
        <div class="input-group w-full">
          <input
            type="text"
            class="input input-bordered w-full"
            :value="icsRoute"
            readonly
          />
          <button class="btn btn-primary" @click.prevent="copyToClipboard()">
            Copy link
          </button>
        </div>
        <a href="#open-how-to-modal" class="btn mt-3 md:ml-3 md:mt-0">
          How to sync?
        </a>
        <HowToModal modalId="open-how-to-modal" />
      </div>
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
          <DocumentOutline class="w-10" />
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
          <NotificationsOutline class="w-10" />
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
      <p v-show="!isDocumentsAvailable" class="h-full">Add any documents</p>
    </div>
  </div>
  <div class="card col-span-2 row-span-4 shadow-lg compact bg-base-100">
    <div class="flex-col items-center card-body">
      <h3 class="card-title text-primary text-left w-full px-4">
        Latest documents
      </h3>
      <ul class="menu flex-row w-full sm:max-h-36 overflow-y-auto">
        <LatestDocumentsItem
          v-for="document in stats.latestDocuments"
          v-bind:key="document.ID"
          v-bind:document="document"
        />
      </ul>
      <p v-show="!isDocumentsAvailable" class="h-full">Add any documents</p>
      <RouterLink to="/documents" class="btn xl:btn-sm w-full btn-primary">
        View all documents
      </RouterLink>
    </div>
  </div>
  <div class="card col-span-3 row-span-1 shadow-lg compact bg-base-100">
    <div class="flex-col sm:flex-row items-center card-body">
      <p v-show="errorMsgs">Export Validity calendar .ics file</p>
      <span
        v-for:="error in errorMsgs"
        class="badge badge-error badge-outline w-full"
      >
        Error: {{ error }}
      </span>
      <button
        v-show="errorMsgs"
        @click.prevent="getIcs"
        class="btn btn-primary btn-sm"
      >
        Export calendar
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, defineAsyncComponent } from "vue";
import { DashboardService } from "./DashboardService";
import type { IDashboardStats } from "./interfaces/IDashboardStats";
import { ErrorDecoder } from "@/services/ErrorDecoder";
import { QueryMaker } from "@/services/QueryMaker";
import { state } from "@/state";

interface VueData {
  stats: IDashboardStats;
  isDocumentsAvailable: boolean;
  avgNotifications: number;
  calendarId: string | null;
  icsRoute: string;
  errorMsgs: string[];
}

export default defineComponent({
  components: {
    HowToModal: defineAsyncComponent(
      () => import("@/components/dashboard/HowToModal.vue")
    ),
  },
  data(): VueData {
    return {
      stats: {} as IDashboardStats,
      isDocumentsAvailable: false,
      avgNotifications: 0,
      calendarId: state.value.user.calendarId,
      icsRoute: "",
      errorMsgs: [],
    };
  },
  methods: {
    async refresh() {
      this.errorMsgs = [];
      this.icsRoute = new QueryMaker({
        route: `/ics/${this.calendarId}`,
      }).routeUrl;
      try {
        this.stats = await DashboardService.getStats();
        this.avgNotifications =
          this.stats.totalNotifications / this.stats.totalDocuments || 0;

        this.isDocumentsAvailable = this.stats.totalDocuments > 0;
      } catch (error) {
        this.errorMsgs.push(await ErrorDecoder.decode(error));
      }
    },
    async getIcs() {
      this.errorMsgs = [];
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
        this.errorMsgs.push(await ErrorDecoder.decode(error));
      }
    },
    copyToClipboard() {
      navigator.clipboard.writeText(this.icsRoute);
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
