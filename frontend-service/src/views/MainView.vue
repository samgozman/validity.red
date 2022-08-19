<script setup lang="ts">
import { RouterLink } from "vue-router";
import MainPageAppHero from "@/components/layout/MainPageAppHero.vue";
import FeaturesRow from "@/components/elements/FeaturesRow.vue";
</script>

<template>
  <main
    class="bg-primary grow min-h-[calc(100%-4rem)] md:min-h-[85%] overflow-hidden"
  >
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
  <div
    class="container sm:mx-auto grow sm:min-h-[calc(100%-12rem)] py-6 sm:py-24"
  >
    <!-- Main features -->
    <div class="pb-6 px-3">
      <span class="font-semibold text-lg text-primary mb-2 block">
        Features
      </span>
      <h2 class="font-bold text-3xl sm:text-4xl md:text-5xl text-dark mb-4">
        Main Features of Validity.RED
      </h2>
      <p
        class="text-lg sm:text-xl leading-relaxed sm:leading-relaxed text-body-color"
      >
        Lorem ipsum dolor sit amet consectetur, adipisicing elit. Temporibus
        alias praesentium quas sint iste tempora modi voluptate placeat. Facere
        accusamus sequi voluptatibus tempore earum aut minus ipsum pariatur ipsa
        eum.
      </p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 gap-6">
      <FeaturesRow
        :rowsData="[
          {
            headline: 'Free',
            text: 'Free. No charges. You can make a donation if you want to support the project.',
            iconStyle: 'heart-outline',
          },
          {
            headline: 'Private',
            text: 'We do not collect any personal information other than your email.',
            iconStyle: 'shield-checkmark-outline',
          },
          {
            headline: 'No extra info',
            text: 'We do not ask you to provide any identifying information. Only the validity period of the document.',
            iconStyle: 'finger-print-outline',
          },
          {
            headline: 'Open-Source',
            text: 'Completely open source. We are not hiding anything from you. In fact, you can join the development!',
            iconStyle: 'code-slash-outline',
          },
        ]"
      />
    </div>
  </div>
</template>
