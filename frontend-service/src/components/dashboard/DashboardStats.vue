<script setup lang="ts">
import UsedTypesItem from "@/components/dashboard/UsedTypesItem.vue";
import LatestDocumentsItem from "@/components/dashboard/LatestDocumentsItem.vue";
</script>

<template>
  <div
    class="card col-span-7 row-span-1 shadow-lg compact bg-base-100 min-h-[40vh]"
  >
    <div class="flex-col sm:flex-row items-center card-body">
      <p>Placeholder for calendar</p>
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
      <p>App version and link to changelog</p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { DashboardService } from "./DashboardService";
import type { IDashboardStats } from "./interfaces/IDashboardStats";

interface VueData {
  stats: IDashboardStats;
  avgNotifications: number;
  error: boolean;
  errorMsg: string;
}

export default defineComponent({
  data(): VueData {
    return {
      stats: {} as IDashboardStats,
      avgNotifications: 0,
      error: false,
      errorMsg: "",
    };
  },
  methods: {
    async refresh() {
      try {
        this.stats = await DashboardService.getStats();
        this.avgNotifications =
          this.stats.totalNotifications / this.stats.totalDocuments || 0;
      } catch (error) {
        this.error = true;
        this.errorMsg = "An error occurred, please try again";
        // TODO: Push error to Sentry
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
