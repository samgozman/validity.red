<script setup lang="ts">
import { RouterLink } from "vue-router";
import MainPageAppHero from "@/components/layout/MainPageAppHero.vue";
</script>

<template>
  <main class="bg-primary grow min-h-[80%]">
    <!-- App representation -->
    <div class="container sm:mx-auto">
      <MainPageAppHero>
        <h1
          class="text-base-100 font-bold text-2xl md:text-3xl lg:text-4xl 2xl:text-5xl md:leading-snug mb-8"
        >
          Keep track on documents validity period. Completely free, private and
          Open-Source.
        </h1>
        <p
          class="text-base-200 text-base sm:text-lg sm:leading-relaxed md:text-xl md:leading-relaxed mx-auto mb-10"
        >
          Manage the validity of documents. Create expiration notifications.
          Synchronize them with your calendar. Always renew on time.
        </p>
        <ul class="flex flex-wrap items-center justify-center mb-10">
          <li>
            <RouterLink
              to="/register"
              class="btn border-none md:px-10 bg-base-100 text-primary hover:shadow-lg hover:bg-primary hover:text-base-100 font-medium transition duration-300 ease-in-out"
            >
              Register
            </RouterLink>
          </li>
          <li>
            <a
              href="https://github.com/samgozman/validity.red"
              target="_blank"
              class="btn border-none hover:bg-primary bg-primary ml-6 font-medium text-base-100 hover:opacity-70 transition duration-300 ease-in-out"
            >
              Star on Github
              <span class="pl-2">
                <ion-icon name="logo-github" size="large"></ion-icon>
              </span>
            </a>
          </li>
        </ul>
      </MainPageAppHero>
    </div>
  </main>
  <div class="grow min-h-[80%]">
    <!-- Main features -->
    <div class="container sm:mx-auto"></div>
  </div>
</template>
