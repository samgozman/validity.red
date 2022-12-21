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
      class="collapse collapse-arrow border border-base-300 bg-base-100 rounded-box"
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
  </ModalBig>
</template>
