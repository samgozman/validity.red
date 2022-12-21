<script setup lang="ts">
import { defineProps } from "vue";
defineProps<{
  modalId: string;
  message: string;
  actionName: string;
}>();
</script>

<template>
  <div v-bind:id="modalId" class="modal modal-bottom sm:modal-middle">
    <div class="modal-box">
      <p>
        {{ message }}
        <strong>{{ actionName }}</strong> ?
      </p>
      <div class="modal-action justify-between sm:justify-end">
        <a href="#" class="btn btn-primary" @click.prevent="clickConfirm"
          >Confirm</a
        >
        <a href="#" class="btn">Close</a>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  methods: {
    clickConfirm() {
      this.$emit("confirmEvent");
    },
  },
});
</script>
