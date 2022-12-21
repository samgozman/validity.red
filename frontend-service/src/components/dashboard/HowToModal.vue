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
      Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil earum
      eos? Eligendi aut, eveniet, quisquam, eius in natus maiores cum quibusdam
      totam pariatur nulla aperiam magni quis praesentium corrupti?
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
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil
          earum eos?
        </p>
        <img
          src="/img/howto/macos_01.webp"
          alt="MacOs calendar - add new Validity.Red subscription"
          class="max-w-xl"
        />
        <p>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil
          earum eos?
        </p>
        <img
          src="/img/howto/macos_02.webp"
          alt="MacOs calendar - edit subscription to Validity.Red"
          class="max-w-xl"
        />
        <p>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil
          earum eos?
        </p>
        <img
          src="/img/howto/macos_03.webp"
          alt="MacOs calendar - view Validity.Red notifications"
          class="max-w-xl"
        />
        <p>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil
          earum eos?
        </p>
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
        <p>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil
          earum eos?
        </p>
        <img
          src="/img/howto/ios_01.jpeg"
          alt="IOS calendar - calendar"
          class="max-w-[16rem]"
        />
        <p>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil
          earum eos?
        </p>
        <img
          src="/img/howto/ios_02.jpeg"
          alt="IOS calendar - calendars list  view"
          class="max-w-[16rem]"
        />
        <p>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil
          earum eos?
        </p>
        <img
          src="/img/howto/ios_03.jpeg"
          alt="IOS calendar - add new Validity.Red subscription"
          class="max-w-[16rem]"
        />
        <p>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil
          earum eos?
        </p>
        <img
          src="/img/howto/ios_04.jpeg"
          alt="IOS calendar - add new Validity.Red subscription"
          class="max-w-[16rem]"
        />
        <p>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil
          earum eos?
        </p>
        <img
          src="/img/howto/ios_05.jpeg"
          alt="IOS calendar - edit Validity.Red subscription options"
          class="max-w-[16rem]"
        />
        <p>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil
          earum eos?
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
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil
          earum eos?
        </p>
        <img
          src="/img/howto/google_01.webp"
          alt="Google calendar"
          class="max-w-xl"
        />
        <p>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil
          earum eos?
        </p>
        <img
          src="/img/howto/google_02.webp"
          alt="Google calendar - subscription to Validity.Red from url"
          class="max-w-xl"
        />
        <p>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil
          earum eos?
        </p>
        <img
          src="/img/howto/google_03.webp"
          alt="Google calendar - add Validity.Red subscription"
          class="max-w-xl"
        />
        <p>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil
          earum eos?
        </p>
        <img
          src="/img/howto/google_04.webp"
          alt="Google calendar - edit Validity.Red notifications"
          class="max-w-xl"
        />
        <p>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde nihil
          earum eos?
        </p>
      </div>
    </div>
  </ModalBig>
</template>
