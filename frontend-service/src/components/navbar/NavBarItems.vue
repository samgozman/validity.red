<script setup lang="ts">
import { RouterLink } from "vue-router";
import { defineProps } from "vue";
import type { NavbarItem } from "@/components/navbar/interfaces/NavbarItem";

defineProps<{
  items: NavbarItem[];
}>();
</script>

<template>
  <li>
    <RouterLink v-for="item in items" v-bind:key="item.name" :to="item.url">
      <ion-icon :name="item.iconClass"></ion-icon> {{ item.name }}
    </RouterLink>
  </li>
</template>
