<script setup lang="ts">
import ModalBig from "@/components/layout/ModalBig.vue";
defineProps<{
  modalId: string;
}>();
</script>

<template>
  <ModalBig :modalId="modalId">
    <h3 class="font-bold text-lg">
      How to sync notifications with your calendar
    </h3>
    <p class="py-4">
      On this tip, you can see various examples of calendar subscriptions. The
      easiest and most popular way to subscribe to Validity.Red is a Google
      Calendar. If all your devices synchronize their calendars through Google,
      then after subscribing to Validity.Red, notifications about document
      expiration dates will appear on all your devices.
    </p>
    <!-- MacOs how-to -->
    <div
      class="collapse collapse-arrow border border-base-300 bg-base-100 rounded-box mb-3"
    >
      <input type="checkbox" />
      <div class="collapse-title text-xl font-medium">
        Sync with macOS Calendar
      </div>
      <div class="collapse-content">
        <p>
          Press
          <kbd class="kbd">⌥</kbd> + <kbd class="kbd">⌘</kbd> +
          <kbd class="kbd">S</kbd>
          to open the "New subscription" window, or click
          <i>File -> New Calendar subscription</i> in the upper menu. <br />
          Put your personal calendar link into the field and press Subscribe.
        </p>
        <img
          src="/img/howto/macos_01.webp"
          alt="MacOs calendar - add new Validity.Red subscription"
          class="w-full max-w-xl"
        />
        <p>
          On the next view, you can set the calendar name, its color, and
          auto-refresh rate.
        </p>
        <img
          src="/img/howto/macos_02.webp"
          alt="MacOs calendar - edit subscription to Validity.Red"
          class="w-full max-w-xl"
        />
        <p>
          That's it! Now you can view your notifications directly from your
          calendar. <br />
          If it doesn't happen for you then you need to refresh your calendar by
          pressing: <kbd class="kbd kbd-sm">⌘</kbd> +
          <kbd class="kbd kbd-sm">R</kbd>.
        </p>
        <img
          src="/img/howto/macos_03.webp"
          alt="MacOs calendar - view Validity.Red notifications"
          class="w-full max-w-xl"
        />
      </div>
    </div>

    <!-- IOS how-to -->
    <div
      class="collapse collapse-arrow border border-base-300 bg-base-100 rounded-box mb-3"
    >
      <input type="checkbox" />
      <div class="collapse-title text-xl font-medium">
        Sync with IOS Calendar
      </div>
      <div class="collapse-content">
        <p class="py-4">
          Open the IOS Calendar app and press the "Calendars" button at the
          bottom of the screen.
        </p>
        <img
          src="/img/howto/ios_01.jpeg"
          alt="IOS calendar - calendar"
          class="w-full max-w-[16rem]"
        />
        <p class="py-4">Then press the "Add Calendar" button at the bottom.</p>
        <img
          src="/img/howto/ios_02.jpeg"
          alt="IOS calendar - calendars list  view"
          class="w-full max-w-[16rem]"
        />
        <p class="py-4">
          In the dropdown menu choose "Add Subscription Calendar".
        </p>
        <img
          src="/img/howto/ios_03.jpeg"
          alt="IOS calendar - add new Validity.Red subscription"
          class="w-full max-w-[16rem]"
        />
        <p class="py-4">
          Put your personal calendar link into the field and press Subscribe.
        </p>
        <img
          src="/img/howto/ios_04.jpeg"
          alt="IOS calendar - add new Validity.Red subscription"
          class="w-full max-w-[16rem]"
        />
        <p class="py-4">
          On the next view, you can set the calendar name, its color, and
          auto-refresh rate.
        </p>
        <img
          src="/img/howto/ios_05.jpeg"
          alt="IOS calendar - edit Validity.Red subscription options"
          class="w-full max-w-[16rem]"
        />
        <p class="pt-4">
          That's it! Now you can view your notifications directly from your
          calendar.
        </p>
      </div>
    </div>

    <!-- Google Calendar how-to -->
    <div
      class="collapse collapse-arrow border border-base-300 bg-base-100 rounded-box mb-3"
    >
      <input type="checkbox" />
      <div class="collapse-title text-xl font-medium">
        Sync with Google Calendar
      </div>
      <div class="collapse-content">
        <p>
          Open
          <a
            href="https://calendar.google.com/calendar"
            target="_blank"
            rel="noopener"
            class="underline text-blue-600 hover:text-blue-800"
          >
            Google Calendar
          </a>
          website. Click on the plus button near the "Other calendars" section
          in the left corner.
        </p>
        <img
          src="/img/howto/google_01.webp"
          alt="Google calendar"
          class="w-full max-w-xl"
        />
        <p>In the dropdown menu choose "From URL"</p>
        <img
          src="/img/howto/google_02.webp"
          alt="Google calendar - subscription to Validity.Red from url"
          class="w-full max-w-xl"
        />
        <p>
          Put your personal calendar link into the field and press Add calendar.
        </p>
        <img
          src="/img/howto/google_03.webp"
          alt="Google calendar - add Validity.Red subscription"
          class="w-full max-w-xl"
        />
        <p>
          On the next view, you can set the calendar name, its color, and
          auto-refresh rate.
        </p>
        <img
          src="/img/howto/google_04.webp"
          alt="Google calendar - edit Validity.Red notifications"
          class="w-full max-w-xl"
        />
        <p>
          That's it! Now you can view your notifications directly from your
          Google calendar.
        </p>
      </div>
    </div>
  </ModalBig>
</template>
