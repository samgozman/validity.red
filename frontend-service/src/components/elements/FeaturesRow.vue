<script setup lang="ts">
import { defineProps } from "vue";

interface RowData {
  headline: string;
  text: string;
  iconStyle: string;
}

defineProps<{
  rowsData: RowData[];
}>();
</script>

<template>
  <div v-for:="row in rowsData" class="card shadow-lg bg-base-100 mx-3 sm:mx-0">
    <div class="card-body">
      <h2 class="my-4 text-xl font-bold card-title">
        <ion-icon :name="row.iconStyle"></ion-icon>
        {{ row.headline }}
      </h2>
      <p class="text-left my-5">{{ row.text }}</p>
    </div>
  </div>
</template>
