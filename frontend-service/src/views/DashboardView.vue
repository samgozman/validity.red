<script setup lang="ts">
import StandardContainer from "@/components/layout/StandardContainer.vue";
import DashboardStats from "@/components/dashboard/DashboardStats.vue";
</script>

<template>
  <StandardContainer>
    <div
      class="grid grid-cols-1 gap-y-6 xl:gap-x-6 py-6 lg:p-10 xl:grid-cols-7 lg:bg-base-200 rounded-box"
    >
      <DashboardStats />
    </div>
  </StandardContainer>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  created() {
    document.title = "Dashboard | Validity";
  },
});
</script>
