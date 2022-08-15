<template>
  <footer class="footer footer-center p-10 bg-primary text-primary-content">
    <div>
      <p class="font-bold">Validity.RED</p>
      <p>Copyright © 2022 - All right reserved</p>
    </div>
    <div>
      <div class="grid grid-flow-col gap-4">
        <a v-for:="button in socialButtons" :href="button.link" target="_blank">
          <ion-icon :name="button.icon" class="text-2xl"></ion-icon>
        </a>
      </div>
    </div>
  </footer>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  data() {
    return {
      socialButtons: [
        {
          link: "https://github.com/samgozman/validity.red",
          icon: "logo-github",
        },
        {
          link: "https://www.linkedin.com/in/samgozman/",
          icon: "logo-linkedin",
        },
        { link: "https://paypal.me/sgozman", icon: "logo-paypal" },
        { link: "https://ko-fi.com/samgozman", icon: "beer-outline" },
      ] as {
        icon: string;
        link: string;
      }[],
    };
  },
});
</script>
