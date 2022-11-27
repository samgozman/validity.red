<script setup lang="ts">
import { RouterLink } from "vue-router";
import { defineProps } from "vue";
import type { NavbarItem } from "@/components/navbar/interfaces/NavbarItem";

defineProps<{
  item: NavbarItem;
  isMobile?: boolean;
}>();
</script>

<template>
  <li>
    <RouterLink :to="item.url" @click.prevent="closeDrawer">
      <component :is="item.icon" class="w-4"></component> {{ item.name }}
    </RouterLink>
  </li>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  methods: {
    closeDrawer() {
      if (!this.isMobile) return;
      document.getElementById("left-sidebar")?.click();
    },
  },
});
</script>
