<script setup lang="ts">
import { RouterLink } from "vue-router";
import MainPageAppHero from "@/components/layout/MainPageAppHero.vue";
import FeaturesRow from "@/components/elements/FeaturesRow.vue";
import LeadingParagraph from "@/components/layout/LeadingParagraph.vue";
import ContrastButton from "@/components/elements/ContrastButton.vue";
import {
  LogoGithub,
  FileTrayFullOutline,
  AlarmOutline,
  CalendarOutline,
  SpeedometerOutline,
  HeartOutline,
  ShieldCheckmarkOutline,
  FingerPrintOutline,
  CodeSlashOutline,
} from "@vicons/ionicons5";
import { state } from "@/state";
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
          Keep track of the document's validity period. Completely free, private
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
            <ContrastButton
              v-if="state.user.isAuthenticated"
              location="/dashboard"
              name="Dashboard"
            />
            <ContrastButton v-else location="/register" name="Register" />
          </li>
          <li>
            <a
              href="https://github.com/samgozman/validity.red"
              target="_blank"
              rel="noopener"
              class="btn border-none hover:bg-primary bg-primary ml-6 font-medium text-base-100 hover:opacity-70 transition duration-300 ease-in-out"
            >
              Star on Github
              <span class="pl-2">
                <LogoGithub class="w-8" />
              </span>
            </a>
          </li>
        </ul>
      </MainPageAppHero>
    </div>
  </main>
  <div class="container sm:mx-auto grow py-6 sm:py-24">
    <!-- Main features -->
    <div class="pb-6 px-3">
      <span class="font-semibold text-lg text-primary mb-2 block">
        Features
      </span>
      <h2 class="font-bold text-3xl sm:text-4xl md:text-5xl text-dark mb-4">
        Main Features of Validity.RED
      </h2>

      <br />
      <LeadingParagraph class="xl:max-w-[60%]">
        Track the expiration dates of your documents with ease in a simple and
        intuitive interface! <br />
        Create multiple notifications for each document. Sync reminders with
        your calendar and never forget to update your documents on time!
      </LeadingParagraph>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 gap-6">
      <FeaturesRow
        :rowsData="[
          {
            headline: 'Manage documents',
            text: 'Track the expiration dates of documents in one place. A large selection of document categories!',
            icon: FileTrayFullOutline,
          },
          {
            headline: 'Create notifications',
            text: `Create multiple notifications to make sure you don't miss it. It's easy and handy!`,
            icon: AlarmOutline,
          },
          {
            headline: 'Sync with calendar',
            text: 'Our most favorite feature: Sync notifications with your calendar. Track reminders in the familiar UI.',
            icon: CalendarOutline,
          },
          {
            headline: 'Handy dashboard',
            text: 'See upcoming expiration dates, statistics and more on a easy-to-use dashboard.',
            icon: SpeedometerOutline,
          },
        ]"
      />
    </div>

    <div class="pt-12 px-3">
      <h2 class="font-bold text-xl sm:text-2xl md:text-3xl text-dark mb-4">
        What Makes us Better
      </h2>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 gap-6">
      <FeaturesRow
        :rowsData="[
          {
            headline: 'Free',
            text: 'Free. No charges. You can make a donation if you want to support the project.',
            icon: HeartOutline,
          },
          {
            headline: 'Private',
            text: 'Data is encrypted with AES-256 algorithm. We do not collect any personal information other than your email.',
            icon: ShieldCheckmarkOutline,
          },
          {
            headline: 'No extra info',
            text: 'We do not ask you to provide any identifying information. Only the validity period of the document.',
            icon: FingerPrintOutline,
          },
          {
            headline: 'Open-Source',
            text: 'Completely open source. We are not hiding anything from you. In fact, you can join the development!',
            icon: CodeSlashOutline,
          },
        ]"
      />
    </div>

    <div class="grid grid-cols-1 py-6 px-3 sm:px-0 sm:pt-24">
      <div class="card shadow-lg bg-base-100">
        <div class="card-body items-center text-center justify-center">
          <div class="xl:max-w-[60%]">
            <h3 class="text-2xl sm:text-3xl md:text-4xl text-dark font-bold">
              Start managing expiration dates!
            </h3>
            <LeadingParagraph class="py-6">
              Has it ever happened to you that when you were going abroad, and
              your visa ran out? You decided to rent a car, but your driver's
              license ran out a week ago? Or your health insurance ended at the
              most inopportune moment and your bank card also expired as well?
              And all of a sudden!
            </LeadingParagraph>
            <LeadingParagraph class="pb-6">
              Sounds like a nightmare! Something similar has happened to all of
              us.
            </LeadingParagraph>
            <LeadingParagraph class="pb-6">
              Join the <strong>Validity.Red</strong> and never forget to update
              documents on time again!
            </LeadingParagraph>
            <RouterLink to="/register" class="btn btn-primary btn-wide btn-lg">
              Get Started
            </RouterLink>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  created() {
    document.title =
      "Validity.Red - manage expiration dates for your documents";
  },
});
</script>
