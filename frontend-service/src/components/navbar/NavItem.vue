<script setup lang="ts">
import { RouterLink } from "vue-router";
import { defineProps } from "vue";
import type { NavbarItem } from "@/components/navbar/interfaces/NavbarItem";

defineProps<{
  item: NavbarItem;
}>();
</script>

<template>
  <li>
    <RouterLink :to="item.url">
      <ion-icon :name="item.iconClass"></ion-icon> {{ item.name }}
    </RouterLink>
  </li>
</template>
