<script setup lang="ts">
import { RouterLink } from "vue-router";
// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { navbarItems } from "@/components/navbar/items";
</script>

<template>
  <footer class="footer footer-center p-10 bg-primary text-primary-content">
    <div class="grid sm:grid-flow-col gap-4">
      <RouterLink
        v-for:="item in navbarItems"
        :to="item.url"
        class="link link-hover"
      >
        <component :is="item.icon" class="w-3 inline"></component>
        {{ item.name }}
      </RouterLink>
    </div>
    <div>
      <p class="font-bold">Validity.RED</p>
      <p>Copyright © 2022 - All right reserved</p>
    </div>
    <div>
      <div class="grid grid-flow-col gap-4">
        <a
          v-for:="button in socialButtons"
          :href="button.link"
          rel="noopener"
          target="_blank"
        >
          <component :is="{ ...button.icon }" class="w-6"></component>
        </a>
      </div>
    </div>
  </footer>
</template>

<script lang="ts">
import type { Component } from "vue";
import { defineComponent } from "vue";
import {
  LogoGithub,
  LogoLinkedin,
  LogoPaypal,
  BeerOutline,
} from "@vicons/ionicons5";

export default defineComponent({
  data() {
    return {
      socialButtons: [
        {
          link: "https://github.com/samgozman/validity.red",
          icon: LogoGithub,
        },
        {
          link: "https://www.linkedin.com/in/samgozman/",
          icon: LogoLinkedin,
        },
        { link: "https://paypal.me/sgozman", icon: LogoPaypal },
        { link: "https://ko-fi.com/samgozman", icon: BeerOutline },
      ] as {
        icon: Component;
        link: string;
      }[],
    };
  },
});
</script>
